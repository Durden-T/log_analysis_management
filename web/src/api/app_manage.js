import service from '@/utils/request'


// @Tags SysApi
// @Summary 创建APP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.App true "创建APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/app [delete]
export const createApp = (data) => {
    return service({
        url: "/app/app",
        method: 'post',
        data
    })
}


// @Tags SysApi
// @Summary 删除APP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.App true "删除APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/app [post]
export const deleteApp = (data) => {
    return service({
        url: "/app/app",
        method: 'delete',
        data
    })
}


// @Tags SysApi
// @Summary 获取权限APP列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取权限APP列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/appList [get]
export const getAppList = (params) => {
    return service({
        url: "/app/appList",
        method: 'get',
        params
    })
}

// @Tags SysApi
// @Summary 更新APP
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.App true "更新APP"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/app [put]
export const updateApp = (data) => {
    return service({
        url: "/app/app",
        method: 'put',
        data
    })
}
