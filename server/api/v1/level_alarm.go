package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags LevelAlarmStrategy
// @Summary 创建LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "创建LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /levelAlarm/createLevelAlarmStrategy [post]
func CreateLevelAlarmStrategy(c *gin.Context) {
	var s model.LevelAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.LevelAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.CreateLevelAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LevelAlarmStrategy
// @Summary 删除LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "删除LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /levelAlarm/deleteLevelAlarmStrategy [delete]
func DeleteLevelAlarmStrategy(c *gin.Context) {
	var s model.LevelAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteLevelAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LevelAlarmStrategy
// @Summary 批量删除LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /levelAlarm/deleteLevelAlarmStrategyByIds [delete]
func DeleteLevelAlarmStrategyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	app := c.GetHeader("app")
	if err := service.DeleteLevelAlarmStrategyByIds(IDS, app); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags LevelAlarmStrategy
// @Summary 更新LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "更新LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /levelAlarm/updateLevelAlarmStrategy [put]
func UpdateLevelAlarmStrategy(c *gin.Context) {
	var s model.LevelAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.LevelAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateLevelAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LevelAlarmStrategy
// @Summary 用id查询LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "用id查询LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /levelAlarm/findLevelAlarmStrategy [get]
func FindLevelAlarmStrategy(c *gin.Context) {
	var s model.LevelAlarmStrategy
	_ = c.ShouldBindQuery(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	app := c.GetHeader("app")
	if err, res := service.GetLevelAlarmStrategy(s.ID, app); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"res": res}, c)
	}
}

// @Tags LevelAlarmStrategy
// @Summary 分页获取LevelAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LevelAlarmStrategySearch true "分页获取LevelAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /levelAlarm/getLevelAlarmStrategyList [get]
func GetLevelAlarmStrategyList(c *gin.Context) {
	var pageInfo request.LevelAlarmStrategySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetLevelAlarmStrategyInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
