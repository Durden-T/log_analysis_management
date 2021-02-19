package model

import (
	"context"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"sort"
	"sync"
	"time"
)

type templateBucket struct {
	bucketSize int
	buckets []map[uint32]*LogTemplate
	result map[uint32]*LogTemplate
	lock sync.RWMutex

	flushInterval time.Duration
	flushEliminationBucketsCount int
	done chan<- struct{}
}

func newTemplateBucket(bucketSize, flushEliminationBucketsCount int, flushInterval time.Duration) *templateBucket{
	bucket := &templateBucket{
		bucketSize: bucketSize,
		flushInterval: flushInterval,
		flushEliminationBucketsCount: flushEliminationBucketsCount,
		buckets: make([]map[uint32]*LogTemplate, bucketSize),
		result: make(map[uint32]*LogTemplate),
	}

	for i := 0; i < bucket.bucketSize; i++ {
		bucket.buckets[i] = make(map[uint32]*LogTemplate, 0)
	}

	bucket.done = utils.SetInterval(bucket.flushInterval, bucket.flush)

	return bucket
}

func (b *templateBucket) flush() error{
	b.lock.Lock()
	defer b.lock.Unlock()

	count := b.flushEliminationBucketsCount
	eliminated := b.buckets[0:count]
	for _, bucket := range eliminated {
		for id, template := range bucket {
			if oldTemplate, found := b.result[id]; found {
				oldTemplate.Size -= template.Size
				if oldTemplate.Size == 0 {
					delete(b.result, id)
				}
			}else {
				fmt.Println("here")
			}
		}
	}

	b.buckets = b.buckets[count:]
	for i:= 0; i < count; i++ {
		b.buckets = append(b.buckets, make(map[uint32]*LogTemplate, 0))
	}
	return nil
}

func(b *templateBucket) Close() {
	b.done<-struct{}{}
}

func(b *templateBucket) Put(index int, t *LogTemplate) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	bucket := b.buckets[index]
	id := t.ClusterId
	if oldTemplate, found := bucket[id]; found {
		oldTemplate.Size += t.Size
	} else {
		bucket[id] = t
	}

	if oldTemplate, found := b.result[id]; found {
		oldTemplate.Size += t.Size
	} else {
		b.result[id] = t.Copy()
	}
}

type TemplateSlice []*LogTemplate

func (s TemplateSlice) Len() int {
	return len(s)
}

func (s TemplateSlice) Less(i, j int) bool {
	return s[i].Size > s[j].Size
}

func (s TemplateSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func(b *templateBucket) GetResult() TemplateSlice {
	b.lock.Lock()
	result := make(TemplateSlice, 0)
	for _, t := range b.result {
		result = append(result, t.Copy())
	}
	b.lock.Unlock()

	sort.Sort(result)
	return result
}

type logParser struct {
	app    string        // app名
	writer *kafka.Writer // 写入日志，生成模版

	reader *kafka.Reader // 获取解析模版的实时结果

	resultBucket *templateBucket // 获取到的结果

	cancel chan struct{}
}

const secondsOfMinute = 60

//@author: [Durden-T](https://github.com/Durden-T)
//@function: NewLogParser
//@description: 创建logParser
//@param: app string, writer *kafka.Writer, reader *kafka.Reader
//@return: *logParser

func NewLogParser(app string, writer *kafka.Writer, reader *kafka.Reader) *logParser {
	return &logParser{
		app:    app,
		reader: reader,
		writer: writer,
		resultBucket: newTemplateBucket(secondsOfMinute, 1, time.Second),
		cancel: make(chan struct{}),
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

	index := time.Now().Second() % secondsOfMinute
	l.resultBucket.Put(index, template)
}

func (l *logParser) FetchResult() TemplateSlice {
	return l.resultBucket.GetResult()
}

func (l *logParser) ProcessLog(log []byte) error {
	return l.writer.WriteMessages(context.TODO(), kafka.Message{
		Value: log,
	})
}
