package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLevelAlarmStrategyRouter(Router *gin.RouterGroup) {
	LevelAlarmStrategyRouter := Router.Group("levelAlarm").Use(middleware.OperationRecord())
	{
		LevelAlarmStrategyRouter.POST("createLevelAlarmStrategy", v1.CreateLevelAlarmStrategy)   // 新建LevelAlarmStrategy
		LevelAlarmStrategyRouter.DELETE("deleteLevelAlarmStrategy", v1.DeleteLevelAlarmStrategy) // 删除LevelAlarmStrategy
		LevelAlarmStrategyRouter.DELETE("deleteLevelAlarmStrategyByIds", v1.DeleteLevelAlarmStrategyByIds) // 批量删除LevelAlarmStrategy
		LevelAlarmStrategyRouter.PUT("updateLevelAlarmStrategy", v1.UpdateLevelAlarmStrategy)    // 更新LevelAlarmStrategy
		LevelAlarmStrategyRouter.GET("findLevelAlarmStrategy", v1.FindLevelAlarmStrategy)        // 根据ID获取LevelAlarmStrategy
		LevelAlarmStrategyRouter.GET("getLevelAlarmStrategyList", v1.GetLevelAlarmStrategyList)  // 获取LevelAlarmStrategy列表
	}
}
