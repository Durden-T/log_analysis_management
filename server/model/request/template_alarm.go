package request

import "gin-vue-admin/model"

type TemplateAlarmStrategySearch struct{
    model.TemplateAlarmStrategy
    PageInfo
}