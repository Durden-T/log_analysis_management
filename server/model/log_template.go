// 自动生成模板LogTemplate
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type LogTemplate struct {
	global.GVA_MODEL
	ClusterId uint32 `json:"clusterId" form:"clusterId" gorm:"column:cluster_id;comment:;uniqueIndex;<-:false;not_null"`
	App       string `json:"app" form:"app" gorm:"column:app;comment:;type:longtext;<-:false;not_null"`
	Tokens    string `json:"tokens" form:"tokens" gorm:"column:tokens;comment:;type:longtext;<-:false;not_null"`
	Size      uint64  `json:"size" form:"size" gorm:"column:size;comment:;type:bigint;<-:false;not_null"`
	Level     string `json:"level" form:"level" gorm:"column:level;comment:;type:longtext;<-:false;not_null"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:;type:longtext;<-:false;not_null"`
}

func (t *LogTemplate) Copy() *LogTemplate {
	return &LogTemplate{
		ClusterId: t.ClusterId,
		App:       t.App,
		Tokens:    t.Tokens,
		Size:      t.Size,
		Level:     t.Level,
		Content:   t.Content,
	}
}
