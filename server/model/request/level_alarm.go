package request

import "gin-vue-admin/model"

type LevelAlarmStrategySearch struct{
    model.LevelAlarmStrategy
    PageInfo
}