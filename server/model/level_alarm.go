// 自动生成模板LevelAlarmStrategy
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type LevelAlarmStrategy struct {
	global.GVA_MODEL
	Interval Duration `json:"interval" form:"interval" gorm:"column:interval;comment:"`
	Count    int64    `json:"count" form:"count" gorm:"column:count;comment:"`
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:;unique"`
	Email    string   `json:"email" form:"email" gorm:"column:email;comment:"`
	App      string   `json:"app" form:"app" gorm:"column:app;comment:;<-:create"`
	Level    string   `json:"level" form:"level" gorm:"column:level;comment:"`

	StartTime  time.Time `json:"-" gorm:"-"`
	StartCount int64     `json:"-" gorm:"-"`
}

const LevelAlarmTableSuffix = "_level_alarms"

func GetLevelAlarmTableName(app string) string {
	return app + LevelAlarmTableSuffix
}
