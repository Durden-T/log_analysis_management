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

// @Tags RegexAlarmStrategy
// @Summary 创建RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "创建RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /regexAlarm/createRegexAlarmStrategy [post]
func CreateRegexAlarmStrategy(c *gin.Context) {
	var s model.RegexAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.RegexAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.CreateRegexAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags RegexAlarmStrategy
// @Summary 删除RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "删除RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /regexAlarm/deleteRegexAlarmStrategy [delete]
func DeleteRegexAlarmStrategy(c *gin.Context) {
	var s model.RegexAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteRegexAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags RegexAlarmStrategy
// @Summary 批量删除RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /regexAlarm/deleteRegexAlarmStrategyByIds [delete]
func DeleteRegexAlarmStrategyByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	app := c.GetHeader("app")
	if err := service.DeleteRegexAlarmStrategyByIds(IDS, app); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags RegexAlarmStrategy
// @Summary 更新RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "更新RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /regexAlarm/updateRegexAlarmStrategy [put]
func UpdateRegexAlarmStrategy(c *gin.Context) {
	var s model.RegexAlarmStrategy
	_ = c.ShouldBindJSON(&s)
	if err := utils.Verify(s, utils.RegexAlarmVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateRegexAlarmStrategy(s); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags RegexAlarmStrategy
// @Summary 用id查询RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "用id查询RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /regexAlarm/findRegexAlarmStrategy [get]
func FindRegexAlarmStrategy(c *gin.Context) {
	var s model.RegexAlarmStrategy
	_ = c.ShouldBindQuery(&s)
	if err := utils.Verify(s.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	app := c.GetHeader("app")
	if err, res := service.GetRegexAlarmStrategy(s.ID, app); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"res": res}, c)
	}
}

// @Tags RegexAlarmStrategy
// @Summary 分页获取RegexAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RegexAlarmStrategySearch true "分页获取RegexAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /regexAlarm/getRegexAlarmStrategyList [get]
func GetRegexAlarmStrategyList(c *gin.Context) {
	var pageInfo request.RegexAlarmStrategySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetRegexAlarmStrategyInfoList(pageInfo); err != nil {
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
