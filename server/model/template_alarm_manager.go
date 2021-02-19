package model

import (
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"github.com/antlabs/timer"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"time"
)

type TemplateAlarmManager struct {
	app    string
	timer.TimeNoder
}

const (
	alarmCheckInterval = time.Second
	TemplateAlarmTableSuffix = "_template_alarms"
)

func NewTemplateAlarmManager(app string) (*TemplateAlarmManager, error) {
	tableName :=  GetTemplateAlarmTableName(app)
	err := global.GVA_DB.Table(tableName).AutoMigrate(&TemplateAlarmStrategy{})
	if err != nil {
		return nil, err
	}
	alarmManager := &TemplateAlarmManager{
		app: app,
	}
	alarmManager.TimeNoder = global.TIMEWHEEL.ScheduleFunc(alarmCheckInterval, alarmManager.checkAlarm)
	return alarmManager, nil
}

func GetTemplateAlarmTableName(app string) string{
	return app + TemplateAlarmTableSuffix
}

func (a *TemplateAlarmManager) checkAlarm() {
	var templateAlarms []*TemplateAlarmStrategy

	db := global.GVA_DB
	if err := db.Table(GetTemplateAlarmTableName(a.app)).Find(&templateAlarms, "app = ?", a.app).Error;
	err != nil {
		global.GVA_LOG.Error("check template alarm failed", zap.String("app", a.app), zap.Error(err))
		return
	}

	if len(templateAlarms) == 0{
		return
	}
	templateIds := make([]uint32, len(templateAlarms))
	for _, alarm := range templateAlarms {
		templateIds = append(templateIds, alarm.TemplateId)
	}
	var templates []*LogTemplate
	if err := db.Table(GetTemplateTableName(a.app)).Find(&templates, "cluster_id in ?", templateIds).Error;
	err != nil {
		global.GVA_LOG.Error("check template alarm failed", zap.String("app", a.app), zap.Error(err))
		return
	}

	now := time.Now()
	for i, alarm := range templateAlarms {
		curCount := templates[i].Size
		if now.Sub(alarm.StartTime) > time.Duration(alarm.Interval) {
			alarm.StartTime = now
			alarm.StartCount = curCount
			continue
		}

		var (
			threshold int64
			increase bool
		)
		if alarm.UseRatio {
			threshold = int64((1+alarm.Ratio)*float64(alarm.StartCount))
			increase = alarm.Ratio > 0
		}else {
			threshold = alarm.StartCount+alarm.Count
			increase = alarm.Count > 0
		}

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
