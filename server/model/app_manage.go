package model

import "gin-vue-admin/global"

type App struct {
	global.GVA_MODEL
	Name       string  `json:"name" form:"name" gorm:"comment:app名, unique, <-:create"`
	KafkaInputTopic string `json:"kafkaInputTopic" form:"kafkaInputTopic" gorm:"comment:kafka输入topic"`
	KafkaOutputTopic string `json:"kafkaOutputTopic" form:"kafkaOutputTopic" gorm:"comment:kafka输出topic"`
	SysUserID          uint    `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
	SysUser            SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`
}

