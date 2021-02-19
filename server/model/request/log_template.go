package request

import (
	"gin-vue-admin/model"
	"regexp"
	"time"
)

type LogTemplateSearch struct {
	model.LogTemplate
	PageInfo
}

type RegexAlarmStrategy struct {
	Interval  time.Duration
	RegexpStr string
	Count     int

	UseRatio   bool
	Ratio      float64
	startCount int
	regex      *regexp.Regexp
}
