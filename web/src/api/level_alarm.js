import service from '@/utils/request'

// @Tags LevelAlarmStrategy
// @Summary 创建LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "创建LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /levelAlarm/createLevelAlarmStrategy [post]
export const createLevelAlarmStrategy = (data) => {
     return service({
         url: "/levelAlarm/createLevelAlarmStrategy",
         method: 'post',
         data
     })
 }


// @Tags LevelAlarmStrategy
// @Summary 删除LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "删除LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /levelAlarm/deleteLevelAlarmStrategy [delete]
 export const deleteLevelAlarmStrategy = (data) => {
     return service({
         url: "/levelAlarm/deleteLevelAlarmStrategy",
         method: 'delete',
         data
     })
 }

// @Tags LevelAlarmStrategy
// @Summary 删除LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /levelAlarm/deleteLevelAlarmStrategy [delete]
 export const deleteLevelAlarmStrategyByIds = (data) => {
     return service({
         url: "/levelAlarm/deleteLevelAlarmStrategyByIds",
         method: 'delete',
         data
     })
 }

// @Tags LevelAlarmStrategy
// @Summary 更新LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "更新LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /levelAlarm/updateLevelAlarmStrategy [put]
 export const updateLevelAlarmStrategy = (data) => {
     return service({
         url: "/levelAlarm/updateLevelAlarmStrategy",
         method: 'put',
         data
     })
 }


// @Tags LevelAlarmStrategy
// @Summary 用id查询LevelAlarmStrategy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LevelAlarmStrategy true "用id查询LevelAlarmStrategy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /levelAlarm/findLevelAlarmStrategy [get]
 export const findLevelAlarmStrategy = (params) => {
     return service({
         url: "/levelAlarm/findLevelAlarmStrategy",
         method: 'get',
         params
     })
 }


// @Tags LevelAlarmStrategy
// @Summary 分页获取LevelAlarmStrategy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LevelAlarmStrategy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /levelAlarm/getLevelAlarmStrategyList [get]
 export const getLevelAlarmStrategyList = (params) => {
     return service({
         url: "/levelAlarm/getLevelAlarmStrategyList",
         method: 'get',
         params
     })
 }