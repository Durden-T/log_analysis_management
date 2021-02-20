package information

import (
	"gin-vue-admin/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"888", 1},
	{"888", 2},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
	{"888", 12},
	{"888", 13},
	{"888", 14},
	{"888", 15},
	{"888", 16},
	{"888", 17},
	{"888", 18},
	{"888", 19},
	{"888", 20},
	{"888", 21},
	{"888", 22},
	{"888", 23},
	{"888", 24},
	{"888", 25},
	{"888", 26},
	{"888", 27},
	{"888", 28},
	{"888", 29},
	{"8881", 1},
	{"8881", 2},
	{"8881", 8},
	{"9528", 1},
	{"9528", 2},
	{"9528", 3},
	{"9528", 4},
	{"9528", 5},
	{"9528", 6},
	{"9528", 7},
	{"9528", 8},
	{"9528", 9},
	{"9528", 10},
	{"9528", 11},
	{"9528", 12},
	{"9528", 14},
	{"9528", 15},
	{"9528", 16},
	{"9528", 17},
	{"10000", 1},
	{"10000", 2},
	{"10000", 3},
	{"10000", 4},
	{"10000", 5},
	{"10000", 6},
	{"10000", 7},
	{"10000", 8},
	{"10000", 9},
	{"10000", 10},
	{"10000", 11},
	{"10000", 12},
	{"10000", 13},
	{"10000", 14},
	{"10000", 15},
	{"10000", 16},
	{"10000", 17},
	{"10000", 18},
	{"10000", 19},
	{"10000", 20},
	{"10000", 21},
	{"10000", 22},
	{"10000", 23},
	{"10000", 24},
	{"10000", 25},
	{"10000", 26},
	{"10000", 27},
	{"10000", 28},
	{"10000", 29},
	{"10000", 10000},
	{"10000", 10001},
	{"10000", 10002},
	{"10000", 10003},
	{"10000", 10004},
	{"10000", 10005},
	{"10000", 10006},
	{"10000", 10007},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.GVA_DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}
