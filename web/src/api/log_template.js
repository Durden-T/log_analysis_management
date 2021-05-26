import service from '@/utils/request'

// @Tags LogTemplate
// @Summary 创建LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogTemplate true "创建LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/createLogTemplate [post]
export const createLogTemplate = (data) => {
     return service({
         url: "/logTemplate/createLogTemplate",
         method: 'post',
         data
     })
 }


// @Tags LogTemplate
// @Summary 删除LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogTemplate true "删除LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logTemplate/deleteLogTemplate [delete]
 export const deleteLogTemplate = (data) => {
     return service({
         url: "/logTemplate/deleteLogTemplate",
         method: 'delete',
         data
     })
 }

// @Tags LogTemplate
// @Summary 删除LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LogTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logTemplate/deleteLogTemplate [delete]
 export const deleteLogTemplateByIds = (data) => {
     return service({
         url: "/logTemplate/deleteLogTemplateByIds",
         method: 'delete',
         data
     })
 }


// @Tags LogTemplate
// @Summary 分页获取LogTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LogTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/getLogTemplateList [get]
 export const getLogTemplateList = (params) => {
     return service({
         url: "/logTemplate/getLogTemplateList",
         method: 'get',
         params
     })
 }


// @Tags LogTemplate
// @Summary 获取实时的LogTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取实时的LogTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logTemplate/getRealtimeResult [get]
 export const getRealtimeResult = (params) => {
     return service({
         url: "/logTemplate/getRealtimeResult",
         method: 'get',
         params
     })
 }

 
// @Tags LogTemplate
// @Summary LogTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LogTemplate true "更新日志模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /logTemplate/updateLogTemplate [put]
export const updateLogTemplate = (data) => {
    return service({
        url: "/logTemplate/updateLogTemplate",
        method: 'put',
        data
    })
}
