package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTemplateAlarmStrategyRouter(Router *gin.RouterGroup) {
	TemplateAlarmStrategyRouter := Router.Group("templateAlarm").Use(middleware.OperationRecord())
	{
		TemplateAlarmStrategyRouter.POST("createTemplateAlarmStrategy", v1.CreateTemplateAlarmStrategy)             // 新建TemplateAlarmStrategy
		TemplateAlarmStrategyRouter.DELETE("deleteTemplateAlarmStrategy", v1.DeleteTemplateAlarmStrategy)           // 删除TemplateAlarmStrategy
		TemplateAlarmStrategyRouter.DELETE("deleteTemplateAlarmStrategyByIds", v1.DeleteTemplateAlarmStrategyByIds) // 批量删除TemplateAlarmStrategy
		TemplateAlarmStrategyRouter.PUT("updateTemplateAlarmStrategy", v1.UpdateTemplateAlarmStrategy)              // 更新TemplateAlarmStrategy
		TemplateAlarmStrategyRouter.GET("findTemplateAlarmStrategy", v1.FindTemplateAlarmStrategy)                  // 根据ID获取TemplateAlarmStrategy
		TemplateAlarmStrategyRouter.GET("getTemplateAlarmStrategyList", v1.GetTemplateAlarmStrategyList)            // 获取TemplateAlarmStrategy列表
	}
}
