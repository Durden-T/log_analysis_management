package model

import (
	"context"
	"gin-vue-admin/global"
	"github.com/antlabs/timer"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gopkg.in/eapache/queue.v1"
	"sort"
	"sync"
	"time"
)

type logTmplMap map[uint32]*LogTemplate

// 时间桶
type templateBucket struct {
	bucketSize int          // 桶的数量
	buckets    *queue.Queue // 元素类型为logTmplMap
	lock       sync.Mutex   // 互斥flush与put

	result sync.Map // key为uint32 模板id, val为*LogTemplate

	cancel timer.TimeNoder // 取消定时淘汰函数
}

func newTemplateBucket(bucketSize int) *templateBucket {
	bucket := &templateBucket{
		bucketSize: bucketSize,
		buckets:    queue.New(),
	}
	bucket.buckets.Add(make(logTmplMap))
	bucket.cancel = global.TIMEWHEEL.ScheduleFunc(time.Second, bucket.flush)

	return bucket
}

func (b *templateBucket) flush() {
	b.lock.Lock()
	b.buckets.Add(make(logTmplMap))
	length := b.buckets.Length()
	if length <= b.bucketSize {
		b.lock.Unlock()
		return
	}

	// 获取最前面的桶
	bucket, ok :=  b.buckets.Remove().(logTmplMap)
	b.lock.Unlock()

	if !ok {
		return
	}
	for id, template := range bucket {
		// 修改result
		if templateInterface, found := b.result.Load(id); found {
			oldTemplate := templateInterface.(*LogTemplate)
			oldTemplate.Size -= template.Size
			if oldTemplate.Size == 0 {
				b.result.Delete(id)
			}
		}
	}
}

func (b *templateBucket) Close() {
	if b.cancel != nil {
		b.cancel.Stop()
	}
}

func (b *templateBucket) Put(t *LogTemplate) {
	b.lock.Lock()
	index := b.buckets.Length() - 1
	// 获取最后一个桶
	bucket, ok := b.buckets.Get(index).(logTmplMap)
	b.lock.Unlock()
	if !ok {
		return
	}
	id := t.ClusterId
	// 已在该index的桶内，增加数量
	if oldTemplate, found := bucket[id]; found {
		oldTemplate.Size += t.Size
	} else {
		// 放入该桶
		bucket[id] = t
	}

	// 处理result: 增加数量/放入result
	if templateInterface, found := b.result.Load(id); found {
		oldTemplate := templateInterface.(*LogTemplate)
		oldTemplate.Size += t.Size
	} else {
		b.result.Store(id, t.Copy())
	}
}

// 为调用sort.sort，实现sort.Interface
type TemplateSlice []*LogTemplate

func (s TemplateSlice) Len() int {
	return len(s)
}

// 根据size降序排序
func (s TemplateSlice) Less(i, j int) bool {
	return s[i].Size > s[j].Size
}

func (s TemplateSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (b *templateBucket) GetResult() TemplateSlice {
	result := make(TemplateSlice, 0)
	b.result.Range(func(key, val interface{}) bool {
		template := val.(*LogTemplate)
		result = append(result, template.Copy())
		return true
	})

	sort.Sort(result)
	return result
}

type logParser struct {
	app    string        // app名
	writer *kafka.Writer // 写入日志，生成模版

	reader *kafka.Reader // 获取解析模版的实时结果

	resultBucket *templateBucket // 获取到的结果

	alarmManager *originalLogAlarmManager // 检查报警

	cancel chan struct{}
}

const secondsOfMinute = 60

//@author: [Durden-T](https://github.com/Durden-T)
//@function: NewLogParser
//@description: 创建logParser
//@param: app string, writer *kafka.Writer, reader *kafka.Reader
//@return: *logParser

func NewLogParser(app string, writer *kafka.Writer, reader *kafka.Reader, alarmManager *originalLogAlarmManager) *logParser {
	return &logParser{
		app:          app,
		reader:       reader,
		writer:       writer,
		resultBucket: newTemplateBucket(secondsOfMinute),
		alarmManager: alarmManager,
		cancel:       make(chan struct{}),
	}
}

func (l *logParser) Run() {
	for {
		select {
		case <-l.cancel:
			return
		default:
			l.updateResult()
		}
	}
}

func (l *logParser) Stop() {
	l.cancel <- struct{}{}
	l.resultBucket.Close()
}

func (l *logParser) updateResult() {
	msg, err := l.reader.ReadMessage(context.TODO())
	if err != nil {
		global.GVA_LOG.Error("read kafka message failed", zap.Any("err", err))
		return
	}

	template := new(LogTemplate)
	if err = jsoniter.Unmarshal(msg.Value, template); err != nil {
		global.GVA_LOG.Error("unmarshal log parser kafka message failed", zap.Any("err", err), zap.Binary("msg.Value", msg.Value))
		return
	}

	// 将日志添加到alarmManager中
	if l.alarmManager != nil {
		l.alarmManager.AddLog(template.Copy())
	}

	l.resultBucket.Put(template)
}

func (l *logParser) FetchResult() TemplateSlice {
	return l.resultBucket.GetResult()
}

// 将原始日志解析成模板
func (l *logParser) ProcessLog(log []byte) error {
	return l.writer.WriteMessages(context.TODO(), kafka.Message{
		Value: log,
	})
}
