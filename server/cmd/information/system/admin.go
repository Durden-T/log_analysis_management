package information

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "ignore", Password: "$2a$10$uLPOjygh/xceuLxF9mKrIO7HuALfHWHwlZi/nMH2xoChao9Hbz7gy", NickName: "ignore", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.jpg", AuthorityId: "888"},
	{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "ignore", Password: "$2a$10$uLPOjygh/xceuLxF9mKrIO7HuALfHWHwlZi/nMH2xoChao9Hbz7gy", NickName: "ignore", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	{GVA_MODEL: global.GVA_MODEL{ID: 10000, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "$2a$10$uLPOjygh/xceuLxF9mKrIO7HuALfHWHwlZi/nMH2xoChao9Hbz7gy", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.png", AuthorityId: "10000"},
	{GVA_MODEL: global.GVA_MODEL{ID: 10001, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "test1", Password: "$2a$10$uLPOjygh/xceuLxF9mKrIO7HuALfHWHwlZi/nMH2xoChao9Hbz7gy", NickName: "测试用户", HeaderImg: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "10001"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
