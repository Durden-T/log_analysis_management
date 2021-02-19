package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLogTemplateRouter(Router *gin.RouterGroup) {
	LogTemplateRouter := Router.Group("logTemplate").Use(middleware.OperationRecord())
	{
		LogTemplateRouter.POST("createLogTemplate", v1.CreateLogTemplate)   // 新建LogTemplate
		LogTemplateRouter.DELETE("deleteLogTemplate", v1.DeleteLogTemplate) // 删除LogTemplate
		LogTemplateRouter.DELETE("deleteLogTemplateByIds", v1.DeleteLogTemplateByIds) // 批量删除LogTemplate
		LogTemplateRouter.GET("getLogTemplateList", v1.GetLogTemplateList)  // 获取LogTemplate列表
		LogTemplateRouter.GET("getRealtimeResult", v1.GetRealtimeResult) //获取实时的日志模版
	}
}
