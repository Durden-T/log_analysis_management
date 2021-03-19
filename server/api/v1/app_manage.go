package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags App
// @Summary 创建App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "App名, kafka输入topic, kafka输出topic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/app [post]
func CreateApp(c *gin.Context) {
	app := new(model.App)
	_ = c.ShouldBindJSON(app)
	if err := utils.Verify(app, utils.AppVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateApp(app); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags App
// @Summary 删除App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "AppID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/app [delete]
func DeleteApp(c *gin.Context) {
	app := new(model.App)
	_ = c.ShouldBindJSON(app)
	if err := utils.Verify(app.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteApp(app); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags App
// @Summary 分页获取权限App列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/appList [get]
func GetAppList(c *gin.Context) {
	pageInfo := new(request.PageInfo)
	_ = c.ShouldBindQuery(pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		pageInfo = nil
	}
	appList, total, err := service.GetAppInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		result := response.PageResult{
			List:  appList,
			Total: total,
		}
		if pageInfo != nil {
			result.Page = pageInfo.Page
			result.PageSize = pageInfo.PageSize
		}
		response.OkWithDetailed(result, "获取成功", c)
	}
}

// @Tags App
// @Summary 更新app, 是否启用报警
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "AppID, 启用报警"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/app [put]
func UpdateApp(c *gin.Context) {
	app := new(model.App)
	_ = c.ShouldBindJSON(app)
	if err := utils.Verify(app.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.UpdateApp(app); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败!", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
