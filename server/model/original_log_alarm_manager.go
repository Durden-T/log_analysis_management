package model

import (
	"gin-vue-admin/global"
	"github.com/antlabs/timer"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type originalLogAlarmManager struct {
	app         string
	levelAlarms []*LevelAlarmStrategy
	regexAlarms []*RegexAlarmStrategy

	buf           []*LogTemplate // 输入日志
	altBuf        []*LogTemplate // 轮换buffer
	switchBufLock sync.RWMutex   // 交换buffer与alternateBuffer
	isChecking    uint32         // 防止多个checkInputLog并发运行

	cancel      chan struct{}
	cancelCheck timer.TimeNoder
}

func NewOriginalLogAlarmManager(app string) (*originalLogAlarmManager, error) {
	err := global.GVA_DB.Table(GetLevelAlarmTableName(app)).AutoMigrate(&LevelAlarmStrategy{})
	if err != nil {
		return nil, err
	}

	err = global.GVA_DB.Table(GetRegexAlarmTableName(app)).AutoMigrate(&RegexAlarmStrategy{})
	if err != nil {
		return nil, err
	}

	manager := &originalLogAlarmManager{
		app:        app,
		cancel:     make(chan struct{}),
		isChecking: 0,
	}
	return manager, nil
}

func (m *originalLogAlarmManager) Start() {
	m.cancelCheck = global.TIMEWHEEL.ScheduleFunc(alarmCheckInterval, m.checkInputLog)
}

func (m *originalLogAlarmManager) Stop() {
	m.cancel <- struct{}{}
	m.cancelCheck.Stop()
}

// 将日志添加到buffer中
func (m *originalLogAlarmManager) AddLog(log *LogTemplate) {
	m.switchBufLock.RLock()
	m.buf = append(m.buf, log)
	m.switchBufLock.RUnlock()
}

func (m *originalLogAlarmManager) checkInputLog() {
	// cas操作 检查此函数是否正在运行
	if !atomic.CompareAndSwapUint32(&m.isChecking, 0, 1) {
		return
	}
	defer func() {
		atomic.StoreUint32(&m.isChecking, 0)
	}()
	// 交换buffer, 减少锁的时间
	m.switchBufLock.Lock()
	buffer := m.buf
	m.buf, m.altBuf = m.altBuf, m.buf
	m.switchBufLock.Unlock()

	if len(buffer) == 0 {
		return
	}

	m.checkLevelAlarm(buffer)
	m.checkRegexAlarm(buffer)
	// 清空备用buffer
	m.altBuf = m.altBuf[:0]
}

func (m *originalLogAlarmManager) checkLevelAlarm(buffer []*LogTemplate) {
	table := GetLevelAlarmTableName(m.app)
	db := global.GVA_DB.Begin()
	defer db.Commit()

	if db.Error != nil {
		return
	}
	if err := m.updateLevelAlarms(db); err != nil {
		global.GVA_LOG.Error("update level alarm failed", zap.Error(err))
	}

	if len(m.levelAlarms) == 0 {
		return
	}

	// 将slice转化为map
	levelAlarms := make(map[string]*LevelAlarmStrategy)
	for _, alarm := range m.levelAlarms {
		levelAlarms[alarm.Level] = alarm
	}

	now := time.Now()
	for _, log := range buffer {
		level := strings.ToLower(log.Level)
		alarm, found := levelAlarms[level]
		// 未配置对应level的报警
		if !found {
			continue
		}
		alarm.CurCount++

		// 与上次检查的间隔大于报警设置的间隔， 跳过本次检查
		if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
			alarm.StartTime = now
			alarm.StartCount = alarm.CurCount
			continue
		}
		if now.Sub(alarm.LastSendTime) < blockAlarmInterval {
			continue
		}
		var (
			threshold uint64
			increase  bool
		)
		if alarm.UseRatio {
			threshold = uint64((1 + alarm.Ratio) * float64(alarm.StartCount))
			increase = alarm.Ratio > 0
		} else {
			threshold = alarm.StartCount + alarm.Count
			increase = alarm.Count > 0
		}
		// 检查报警条件
		if (increase && alarm.CurCount >= threshold) || (!increase && alarm.CurCount <= threshold) {
			if err := sendAlarm(alarm, alarm.Email); err != nil {
				global.GVA_LOG.Error("send level alarm failed", zap.String("app", alarm.App),
					zap.Any("level_alarm", alarm), zap.Error(err))
				continue
			}
			alarm.LastSendTime = now
		}
	}
	// 保存startCount与startTimes
	if err := db.Table(table).Save(&m.levelAlarms).Error; err != nil {
		global.GVA_LOG.Error("save level alarm failed", zap.String("app", m.app), zap.Error(err))
	}
}

func (m *originalLogAlarmManager) checkRegexAlarm(buffer []*LogTemplate) {
	table := GetLevelAlarmTableName(m.app)
	db := global.GVA_DB.Begin()
	defer db.Commit()
	if db.Error != nil {
		return
	}
	if err := m.updateRegexAlarms(db); err != nil {
		global.GVA_LOG.Error("update regex alarm failed", zap.Error(err))
	}

	if len(m.regexAlarms) == 0 {
		return
	}

	now := time.Now()
	for _, log := range buffer {
		for _, alarm := range m.regexAlarms {
			if alarm.Regexp == nil {
				continue
			}
			if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
				alarm.StartTime = now
				alarm.StartCount = alarm.CurCount
				continue
			}
			if now.Sub(alarm.LastSendTime) < blockAlarmInterval {
				continue
			}
			if alarm.Regexp.MatchString(log.Content) {
				alarm.CurCount++
			}

			var (
				threshold uint64
				increase  bool
			)
			if alarm.UseRatio {
				threshold = uint64((1 + alarm.Ratio) * float64(alarm.StartCount))
				increase = alarm.Ratio > 0
			} else {
				threshold = alarm.StartCount + alarm.Count
				increase = alarm.Count > 0
			}
			// 检查报警条件
			if (increase && alarm.CurCount >= threshold) || (!increase && alarm.CurCount <= threshold) {
				if err := sendAlarm(alarm, alarm.Email); err != nil {
					global.GVA_LOG.Error("send regex alarm failed", zap.String("app", alarm.App),
						zap.Any("regex_alarm", alarm), zap.Error(err))
					continue
				}
				alarm.LastSendTime = now
			}
		}
	}
	// 保存startCount与startTimes
	if err := db.Table(table).Save(&m.regexAlarms).Error; err != nil {
		global.GVA_LOG.Error("save regex alarm failed", zap.String("app", m.app), zap.Error(err))
	}
}

func (m *originalLogAlarmManager) updateLevelAlarms(db *gorm.DB) error {
	table := GetLevelAlarmTableName(m.app)
	// 上行锁
	return db.Table(table).Clauses(clause.Locking{Strength: "UPDATE"}).Find(&m.levelAlarms).Error
}

func (m *originalLogAlarmManager) updateRegexAlarms(db *gorm.DB) error {
	table := GetRegexAlarmTableName(m.app)
	// 上行锁
	return db.Table(table).Clauses(clause.Locking{Strength: "UPDATE"}).Find(&m.regexAlarms).Error
}
