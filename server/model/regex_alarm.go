package model

import (
	"gin-vue-admin/global"
	"regexp"
	"time"
)

// 如果含有time.Time 请自行import time包
type RegexAlarmStrategy struct {
	global.GVA_MODEL
	Interval Duration `json:"interval" form:"interval" gorm:"column:interval;comment:"`
	Count    int64    `json:"count" form:"count" gorm:"column:count;comment:"`
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:;unique"`
	Email    string   `json:"email" form:"email" gorm:"column:email;comment:"`
	App      string   `json:"app" form:"app" gorm:"column:app;comment:;<-:create"`
	Regex    string   `json:"regex" form:"regex" gorm:"column:regex;comment:"`

	*regexp.Regexp `json:"-" gorm:"-"`
	StartTime      time.Time `json:"-"`
	StartCount     int64     `json:"-"`
}

const RegexAlarmTableSuffix = "_regex_alarms"

func GetRegexAlarmTableName(app string) string {
	return app + RegexAlarmTableSuffix
}
