package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRegexAlarmStrategyRouter(Router *gin.RouterGroup) {
	RegexAlarmStrategyRouter := Router.Group("regexAlarm").Use(middleware.OperationRecord())
	{
		RegexAlarmStrategyRouter.POST("createRegexAlarmStrategy", v1.CreateRegexAlarmStrategy)   // 新建RegexAlarmStrategy
		RegexAlarmStrategyRouter.DELETE("deleteRegexAlarmStrategy", v1.DeleteRegexAlarmStrategy) // 删除RegexAlarmStrategy
		RegexAlarmStrategyRouter.DELETE("deleteRegexAlarmStrategyByIds", v1.DeleteRegexAlarmStrategyByIds) // 批量删除RegexAlarmStrategy
		RegexAlarmStrategyRouter.PUT("updateRegexAlarmStrategy", v1.UpdateRegexAlarmStrategy)    // 更新RegexAlarmStrategy
		RegexAlarmStrategyRouter.GET("findRegexAlarmStrategy", v1.FindRegexAlarmStrategy)        // 根据ID获取RegexAlarmStrategy
		RegexAlarmStrategyRouter.GET("getRegexAlarmStrategyList", v1.GetRegexAlarmStrategyList)  // 获取RegexAlarmStrategy列表
	}
}
