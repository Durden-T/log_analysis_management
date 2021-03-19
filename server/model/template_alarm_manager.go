package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"github.com/antlabs/timer"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"time"
)

type templateAlarmManager struct {
	app string
	timer.TimeNoder
}

const (
	alarmCheckInterval       = 500 * time.Millisecond // 报警检查间隔
	TemplateAlarmTableSuffix = "_template_alarms"
)

func NewTemplateAlarmManager(app string) (*templateAlarmManager, error) {
	tableName := GetTemplateAlarmTableName(app)
	err := global.GVA_DB.Table(tableName).AutoMigrate(&TemplateAlarmStrategy{})
	if err != nil {
		return nil, err
	}
	alarmManager := &templateAlarmManager{
		app: app,
	}
	return alarmManager, nil
}

func GetTemplateAlarmTableName(app string) string {
	return app + TemplateAlarmTableSuffix
}

func (a *templateAlarmManager) Start() {
	a.TimeNoder = global.TIMEWHEEL.ScheduleFunc(alarmCheckInterval, a.checkAlarm)
}

func (a *templateAlarmManager) checkAlarm() {
	var templateAlarms []*TemplateAlarmStrategy

	db := global.GVA_DB.Begin()
	defer db.Commit()
	// 获取对应app的模板报警信息; 上行锁
	if err := db.Table(GetTemplateAlarmTableName(a.app)).Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&templateAlarms, "app = ?", a.app).Error; err != nil {
		global.GVA_LOG.Error("check template alarm failed", zap.String("app", a.app), zap.Error(err))
		return
	}

	if len(templateAlarms) == 0 {
		return
	}
	// 保存模板报警对应的id信息
	templateIds := make([]uint32, len(templateAlarms))
	for _, alarm := range templateAlarms {
		templateIds = append(templateIds, alarm.TemplateId)
	}
	// 获取对应id的模板
	var templates []*LogTemplate
	if err := db.Table(GetTemplateTableName(a.app)).Find(&templates, "cluster_id in ?", templateIds).
		Error; err != nil {
		global.GVA_LOG.Error("check template alarm failed", zap.String("app", a.app), zap.Error(err))
		return
	}
	// 检查模板对应的数量变化
	now := time.Now()
	for i, alarm := range templateAlarms {
		curCount := templates[i].Size
		// 超过间隔
		if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
			alarm.StartTime = now
			alarm.StartCount = curCount
			continue
		}

		var (
			threshold int64
			increase  bool
		)
		if alarm.UseRatio {
			threshold = int64((1 + alarm.Ratio) * float64(alarm.StartCount))
			increase = alarm.Ratio > 0
		} else {
			threshold = alarm.StartCount + alarm.Count
			increase = alarm.Count > 0
		}
		// 检查报警条件
		if (increase && curCount >= threshold) || (!increase && curCount <= threshold) {
			if err := sendAlarm(alarm, alarm.Email); err != nil {
				global.GVA_LOG.Error("send template alarm failed", zap.String("app", a.app),
					zap.Any("template_alarm", alarm), zap.Error(err))
			}
		}
	}
	if err := db.Table(GetTemplateAlarmTableName(a.app)).Save(&templateAlarms).Error; err != nil {
		global.GVA_LOG.Error("save template alarm failed", zap.String("app", a.app), zap.Error(err))
	}
}

func sendAlarm(data interface{}, email string) error {
	body, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}
	return utils.Send([]string{email}, "日志报警", string(body))
}
