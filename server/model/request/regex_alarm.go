package request

import "gin-vue-admin/model"

type RegexAlarmStrategySearch struct{
    model.RegexAlarmStrategy
    PageInfo
}