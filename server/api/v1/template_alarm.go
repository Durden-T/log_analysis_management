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

// @Tags TemplateAlarmStrategy
// @Summary 创建TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "创建TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /templateAlarm/createTemplateAlarmStrategy [post]
func CreateTemplateAlarmStrategy(c *gin.Context) {
	var s model.TemplateAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.TemplateAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateTemplateAlarmStrategy(s); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TemplateAlarmStrategy
// @Summary 删除TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "删除TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /templateAlarm/deleteTemplateAlarmStrategy [delete]
func DeleteTemplateAlarmStrategy(c *gin.Context) {
	var s model.TemplateAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteTemplateAlarmStrategy(s); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TemplateAlarmStrategy
// @Summary 批量删除TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /templateAlarm/deleteTemplateAlarmStrategyByIds [delete]
func DeleteTemplateAlarmStrategyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	app := c.GetHeader("app")
	if err := service.DeleteTemplateAlarmStrategyByIds(IDS, app); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TemplateAlarmStrategy
// @Summary 更新TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "更新TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /templateAlarm/updateTemplateAlarmStrategy [put]
func UpdateTemplateAlarmStrategy(c *gin.Context) {
	var s model.TemplateAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.TemplateAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateTemplateAlarmStrategy(s); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TemplateAlarmStrategy
// @Summary 用id查询TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "用id查询TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /templateAlarm/findTemplateAlarmStrategy [get]
func FindTemplateAlarmStrategy(c *gin.Context) {
	var s model.TemplateAlarmStrategy
	_ = c.ShouldBindQuery(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := service.GetTemplateAlarmStrategy(s.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"res": res}, c)
	}
}

// @Tags TemplateAlarmStrategy
// @Summary 分页获取TemplateAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TemplateAlarmStrategySearch true "分页获取TemplateAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /templateAlarm/getTemplateAlarmStrategyList [get]
func GetTemplateAlarmStrategyList(c *gin.Context) {
	var pageInfo request.TemplateAlarmStrategySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetTemplateAlarmStrategyInfoList(pageInfo); err != nil {
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
