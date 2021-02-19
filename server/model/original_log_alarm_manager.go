package model

import (
	"context"
	"gin-vue-admin/global"
	"github.com/antlabs/timer"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"regexp"
	"strings"
	"sync"
	"time"
)

type OriginalLogAlarmManager struct {
	app         string
	reader      *kafka.Reader // 获取原始日志
	levelAlarms map[string]*LevelAlarmStrategy
	regexAlarms map[string]*RegexAlarmStrategy
	lock        sync.Mutex

	cancel             chan struct{}
	cancelUpdateAlarms timer.TimeNoder
}

func NewOriginalLogAlarmManager(app string, reader *kafka.Reader) (*OriginalLogAlarmManager, error) {
	err := global.GVA_DB.Table(GetLevelAlarmTableName(app)).AutoMigrate(&LevelAlarmStrategy{})
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.Table(GetRegexAlarmTableName(app)).AutoMigrate(&RegexAlarmStrategy{})
	if err != nil {
		return nil, err
	}

	manager := &OriginalLogAlarmManager{
		app:    app,
		reader: reader,
		cancel: make(chan struct{}),
	}
	manager.cancelUpdateAlarms = global.TIMEWHEEL.ScheduleFunc(5*time.Second, manager.updateAlarms)
	return manager, nil
}

func (m *OriginalLogAlarmManager) Run() {
	for {
		select {
		case <-m.cancel:
			return
		default:
			m.checkInputLog()
		}
	}
}

func (m *OriginalLogAlarmManager) Stop() {
	m.cancel <- struct{}{}
	m.cancelUpdateAlarms.Stop()
}

func (m *OriginalLogAlarmManager) checkInputLog() {
	msg, err := m.reader.ReadMessage(context.TODO())
	if err != nil {
		global.GVA_LOG.Error("read kafka message failed", zap.Any("err", err))
		return
	}
	if len(msg.Value) == 0 {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.checkLevelAlarm(msg.Value)
	m.checkRegexAlarm(msg.Value)
}

func (m *OriginalLogAlarmManager) checkLevelAlarm(val []byte) {
	if len(m.levelAlarms) == 0 {
		return
	}
	level := jsoniter.Get(val, "level").ToString()
	level = strings.ToLower(level)
	now := time.Now()
	if alarm, found := m.levelAlarms[level]; found {
		if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
			alarm.StartTime = now
			alarm.StartCount = 0
			return
		}
		alarm.StartCount++
		if alarm.StartCount >= alarm.Count {
			if err := sendAlarm(alarm, alarm.Email); err != nil {
				global.GVA_LOG.Error("send level alarm failed", zap.String("app", alarm.App),
					zap.Any("level_alarm", alarm), zap.Error(err))
			}
		}
	}
}

func (m *OriginalLogAlarmManager) checkRegexAlarm(val []byte) {
	if len(m.regexAlarms) == 0 {
		return
	}
	content := jsoniter.Get(val, "content").ToString()
	now := time.Now()
	for _, alarm := range m.regexAlarms {
		if alarm.Regexp == nil {
			continue
		}
		if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
			alarm.StartTime = now
			alarm.StartCount = 0
			return
		}

		if alarm.Regexp.MatchString(content) {
			alarm.StartCount++
		}

		if alarm.StartCount >= alarm.Count {
			if err := sendAlarm(alarm, alarm.Email); err != nil {
				global.GVA_LOG.Error("send regex alarm failed", zap.String("app", alarm.App),
					zap.Any("regex_alarm", alarm), zap.Error(err))
			}
		}
	}
}

func (m *OriginalLogAlarmManager) updateAlarms() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.updateLevelAlarms()
	m.updateRegexAlarms()
}

func (m *OriginalLogAlarmManager) updateLevelAlarms() {
	var levelAlarms []*LevelAlarmStrategy
	table := GetLevelAlarmTableName(m.app)
	global.GVA_DB.Table(table).Find(&levelAlarms)
	for _, alarm := range levelAlarms {
		if oldAlarm, found := m.levelAlarms[alarm.Name]; found {
			alarm.StartCount = oldAlarm.StartCount
			alarm.StartTime = oldAlarm.StartTime
		}
	}
	m.levelAlarms = make(map[string]*LevelAlarmStrategy)
	for _, alarm := range levelAlarms {
		m.levelAlarms[alarm.Name] = alarm
	}
}

func (m *OriginalLogAlarmManager) updateRegexAlarms() {
	var regexAlarms []*RegexAlarmStrategy

	table := GetRegexAlarmTableName(m.app)
	global.GVA_DB.Table(table).Find(&regexAlarms)

	for _, alarm := range regexAlarms {
		if oldAlarm, found := m.regexAlarms[alarm.Name]; found {
			alarm.StartCount = oldAlarm.StartCount
			alarm.StartTime = oldAlarm.StartTime
		}
	}
	m.regexAlarms = make(map[string]*RegexAlarmStrategy)
	var err error
	for _, alarm := range regexAlarms {
		if alarm.Regexp, err = regexp.Compile(alarm.Regex); err != nil {
			global.GVA_LOG.Error("compile regex alarm failed", zap.String("app", alarm.App),
				zap.Any("regex_alarm", alarm), zap.Error(err))
			continue
		}
		m.regexAlarms[alarm.Name] = alarm
	}
}
