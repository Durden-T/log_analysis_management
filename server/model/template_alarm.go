// 自动生成模板TemplateAlarmStrategy
package model

import (
	"encoding/json"
	"errors"
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type TemplateAlarmStrategy struct {
	global.GVA_MODEL
	Interval Duration `json:"interval" form:"interval" gorm:"column:interval;comment:"`
	Count    uint64   `json:"count" form:"count" gorm:"column:count;comment:"`
	UseRatio bool     `json:"useRatio" form:"useRatio" gorm:"column:use_ratio;comment:"`
	Ratio    float64  `json:"ratio" form:"ratio" gorm:"column:ratio;comment:"`
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:;unique"`
	Email    string   `json:"email" form:"email" gorm:"column:email;comment:"`
	App      string   `json:"app" form:"app" gorm:"<-:create"`

	TemplateId uint32 `json:"templateId" form:"templateId" gorm:"<-:create"`

	StartTime    time.Time `json:"-"`
	StartCount   uint64    `json:"-"`
	LastSendTime time.Time `json:"-"`
}

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}
