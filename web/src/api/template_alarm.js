import service from '@/utils/request'

// @Tags TemplateAlarmStrategy
// @Summary 创建TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "创建TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /templateAlarm/createTemplateAlarmStrategy [post]
export const createTemplateAlarmStrategy = (data) => {
     return service({
         url: "/templateAlarm/createTemplateAlarmStrategy",
         method: 'post',
         data
     })
 }


// @Tags TemplateAlarmStrategy
// @Summary 删除TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "删除TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /templateAlarm/deleteTemplateAlarmStrategy [delete]
 export const deleteTemplateAlarmStrategy = (data) => {
     return service({
         url: "/templateAlarm/deleteTemplateAlarmStrategy",
         method: 'delete',
         data
     })
 }

// @Tags TemplateAlarmStrategy
// @Summary 删除TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /templateAlarm/deleteTemplateAlarmStrategy [delete]
 export const deleteTemplateAlarmStrategyByIds = (data) => {
     return service({
         url: "/templateAlarm/deleteTemplateAlarmStrategyByIds",
         method: 'delete',
         data
     })
 }

// @Tags TemplateAlarmStrategy
// @Summary 更新TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "更新TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /templateAlarm/updateTemplateAlarmStrategy [put]
 export const updateTemplateAlarmStrategy = (data) => {
     return service({
         url: "/templateAlarm/updateTemplateAlarmStrategy",
         method: 'put',
         data
     })
 }


// @Tags TemplateAlarmStrategy
// @Summary 用id查询TemplateAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TemplateAlarmStrategy true "用id查询TemplateAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /templateAlarm/findTemplateAlarmStrategy [get]
 export const findTemplateAlarmStrategy = (params) => {
     return service({
         url: "/templateAlarm/findTemplateAlarmStrategy",
         method: 'get',
         params
     })
 }


// @Tags TemplateAlarmStrategy
// @Summary 分页获取TemplateAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TemplateAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /templateAlarm/getTemplateAlarmStrategyList [get]
 export const getTemplateAlarmStrategyList = (params) => {
     return service({
         url: "/templateAlarm/getTemplateAlarmStrategyList",
         method: 'get',
         params
     })
 }