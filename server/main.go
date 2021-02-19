package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"github.com/antlabs/timer"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
	global.TIMEWHEEL = timer.NewTimer() // 初始化时间轮
	go global.TIMEWHEEL.Run()
	defer global.TIMEWHEEL.Stop()

	initialize.InitBusiness()

	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	core.RunWindowsServer()
}
