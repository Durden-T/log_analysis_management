package information

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: true, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: model.Meta{Title: "关于我们", Icon: "info"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: model.Meta{Title: "示例文件", Icon: "s-management"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: model.Meta{Title: "excel导入导出", Icon: "s-marketing"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: model.Meta{Title: "上传下载", Icon: "upload"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: model.Meta{Title: "断点续传", Icon: "upload"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: model.Meta{Title: "客户列表（资源示例）", Icon: "s-custom"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: model.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: model.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true}},
	{GVA_MODEL: global.GVA_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: model.Meta{Title: "断点续传（插件版）", Icon: "upload"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Hidden: true, Component: "/", Sort: 0, Meta: model.Meta{Title: "官方网站", Icon: "s-home"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "workflow", Name: "workflow", Hidden: true, Component: "view/workflow/index.vue", Sort: 5, Meta: model.Meta{Title: "工作流功能", Icon: "phone"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowCreate", Name: "workflowCreate", Hidden: false, Component: "view/workflow/workflowCreate/workflowCreate.vue", Sort: 0, Meta: model.Meta{Title: "工作流绘制", Icon: "circle-plus"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowProcess", Name: "workflowProcess", Hidden: false, Component: "view/workflow/workflowProcess/workflowProcess.vue", Sort: 0, Meta: model.Meta{Title: "工作流列表", Icon: "s-cooperation"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "workflowUse", Name: "workflowUse", Hidden: true, Component: "view/workflow/workflowUse/workflowUse.vue", Sort: 0, Meta: model.Meta{Title: "使用工作流", Icon: "video-play"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "started", Name: "started", Hidden: false, Component: "view/workflow/userList/started.vue", Sort: 0, Meta: model.Meta{Title: "我发起的", Icon: "s-order"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "24", Path: "need", Name: "need", Hidden: false, Component: "view/workflow/userList/need.vue", Sort: 0, Meta: model.Meta{Title: "我的待办", Icon: "s-platform"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10000, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "app", Name: "app", Hidden: false, Component: "view/app_manage/app_manage.vue", Sort: 0, Meta: model.Meta{Title: "app管理", Icon: "s-management"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10001, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "logTemplate", Name: "logTemplate", Hidden: false, Component: "view/log_template/index.vue", Sort: 1, Meta: model.Meta{Title: "日志模版", Icon: "s-order"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10002, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "10001", Path: "search", Name: "search", Hidden: false, Component: "view/log_template/template_search/log_template.vue", Sort: 0, Meta: model.Meta{Title: "模版搜索", Icon: "search"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10003, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "10001", Path: "realtime", Name: "realtime", Hidden: false, Component: "view/log_template/realtime/realtime.vue", Sort: 1, Meta: model.Meta{Title: "实时模版", Icon: "info"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10004, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "alarm", Name: "alarm", Hidden: false, Component: "view/alarm/index.vue", Sort: 2, Meta: model.Meta{Title: "报警策略", Icon: "alarm-clock"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10005, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "10004", Path: "templateAlarm", Name: "templateAlarm", Hidden: false, Component: "view/alarm/template_alarm/template_alarm.vue", Sort: 0, Meta: model.Meta{Title: "模版报警", Icon: "s-goods"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10006, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "10004", Path: "levelAlarm", Name: "levelAlarm", Hidden: false, Component: "view/alarm/level_alarm/level_alarm.vue", Sort: 1, Meta: model.Meta{Title: "级别报警", Icon: "files"}},
	{GVA_MODEL: global.GVA_MODEL{ID: 10007, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "10004", Path: "regexAlarm", Name: "regexAlarm", Hidden: false, Component: "view/alarm/regex_alarm/regex_alarm.vue", Sort: 1, Meta: model.Meta{Title: "正则报警", Icon: "notebook-1"}},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
