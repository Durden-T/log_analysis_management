package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAppRouter(Router *gin.RouterGroup) {
	AppRouter := Router.Group("app").Use(middleware.OperationRecord())
	{
		AppRouter.POST("app", v1.CreateApp)     // 创建App
		AppRouter.DELETE("app", v1.DeleteApp)   // 删除App
		AppRouter.GET("appList", v1.GetAppList) // 获取App列表
		AppRouter.PUT("app", v1.UpdateApp) // 更新app是否启用报警
	}
}
