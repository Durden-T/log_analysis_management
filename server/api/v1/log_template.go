package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags LogTemplate
// @Summary 创建LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogTemplate true "创建LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/createLogTemplate [post]
func CreateLogTemplate(c *gin.Context) {
	var l model.LogTemplate
	_ = c.ShouldBindJSON(&l)
	if err := service.CreateLogTemplate(l); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LogTemplate
// @Summary 删除LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogTemplate true "删除LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logTemplate/deleteLogTemplate [delete]
func DeleteLogTemplate(c *gin.Context) {
	var l model.LogTemplate
	_ = c.ShouldBindJSON(&l)
	if err := service.DeleteLogTemplate(l); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LogTemplate
// @Summary 批量删除LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /logTemplate/deleteLogTemplateByIds [delete]
func DeleteLogTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	app := c.GetHeader("app")
	if err := service.DeleteLogTemplateByIds(IDS, app); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags LogTemplate
// @Summary 分页获取LogTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LogTemplateSearch true "分页获取LogTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/getLogTemplateList [get]
func GetLogTemplateList(c *gin.Context) {
	var pageInfo request.LogTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLogTemplateInfoList(pageInfo); err != nil {
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


// @Tags LogTemplate
// @Summary 获取实时的日志模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/getRealtimeResult [get]
func GetRealtimeResult(c *gin.Context) {
	app := c.GetHeader("app")
	oldAppInterface, ok := global.APP_MANAGER.Load(app)
	if !ok {
		response.FailWithMessage("获取失败", c)
		return
	}
	oldApp, ok := oldAppInterface.(*model.App)
	if !ok {
		response.FailWithMessage("获取失败", c)
		return

	}
	logParser := oldApp.LogParser
	res := logParser.FetchResult()
	if len(res) == 0 {
		global.GVA_LOG.Error("获取失败")
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     res,
		}, "获取成功", c)
	}
}
