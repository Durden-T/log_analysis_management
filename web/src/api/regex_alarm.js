import service from '@/utils/request'

// @Tags RegexAlarmStrategy
// @Summary 创建RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "创建RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /regexAlarm/createRegexAlarmStrategy [post]
export const createRegexAlarmStrategy = (data) => {
     return service({
         url: "/regexAlarm/createRegexAlarmStrategy",
         method: 'post',
         data
     })
 }


// @Tags RegexAlarmStrategy
// @Summary 删除RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "删除RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /regexAlarm/deleteRegexAlarmStrategy [delete]
 export const deleteRegexAlarmStrategy = (data) => {
     return service({
         url: "/regexAlarm/deleteRegexAlarmStrategy",
         method: 'delete',
         data
     })
 }

// @Tags RegexAlarmStrategy
// @Summary 删除RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /regexAlarm/deleteRegexAlarmStrategy [delete]
 export const deleteRegexAlarmStrategyByIds = (data) => {
     return service({
         url: "/regexAlarm/deleteRegexAlarmStrategyByIds",
         method: 'delete',
         data
     })
 }

// @Tags RegexAlarmStrategy
// @Summary 更新RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "更新RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /regexAlarm/updateRegexAlarmStrategy [put]
 export const updateRegexAlarmStrategy = (data) => {
     return service({
         url: "/regexAlarm/updateRegexAlarmStrategy",
         method: 'put',
         data
     })
 }


// @Tags RegexAlarmStrategy
// @Summary 用id查询RegexAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RegexAlarmStrategy true "用id查询RegexAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /regexAlarm/findRegexAlarmStrategy [get]
 export const findRegexAlarmStrategy = (params) => {
     return service({
         url: "/regexAlarm/findRegexAlarmStrategy",
         method: 'get',
         params
     })
 }


// @Tags RegexAlarmStrategy
// @Summary 分页获取RegexAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RegexAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /regexAlarm/getRegexAlarmStrategyList [get]
 export const getRegexAlarmStrategyList = (params) => {
     return service({
         url: "/regexAlarm/getRegexAlarmStrategyList",
         method: 'get',
         params
     })
 }