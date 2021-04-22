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
	Count    uint64    `json:"count" form:"count" gorm:"column:count;comment:"`
	UseRatio bool     `json:"useRatio" form:"useRatio" gorm:"column:use_ratio;comment:"`
	Ratio    float64  `json:"ratio" form:"ratio" gorm:"column:ratio;comment:"`
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:;unique"`
	Email    string   `json:"email" form:"email" gorm:"column:email;comment:"`
	App      string   `json:"app" form:"app" gorm:"column:app;comment:;<-:create"`
	Regex    string   `json:"regex" form:"regex" gorm:"column:regex;comment:"`

	*regexp.Regexp `json:"-" gorm:"-"`
	CurCount   uint64     `json:"-"`
	StartTime      time.Time `json:"-"`
	StartCount     uint64     `json:"-"`
	LastSendTime time.Time `json:"-"`
}

const RegexAlarmTableSuffix = "_regex_alarms"

func GetRegexAlarmTableName(app string) string {
	return app + RegexAlarmTableSuffix
}
