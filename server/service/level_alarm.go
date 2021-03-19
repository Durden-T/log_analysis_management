package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"
	"time"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: CreateLevelAlarmStrategy
//@description: 创建LevelAlarmStrategy记录
//@param: s model.LevelAlarmStrategy
//@return: err error

func CreateLevelAlarmStrategy(s model.LevelAlarmStrategy) (err error) {
	if _, found := global.APP_MANAGER.Load(s.App); !found {
		return errors.New("app doesn't exist")
	}
	s.Level = strings.ToLower(s.Level)

	if checkLevelExist(s.App, s.Level) {
		return errors.New("level existed")
	}

	now := time.Now()
	s.StartTime = now
	s.StartCount = 0
	err = global.GVA_DB.Table(model.GetLevelAlarmTableName(s.App)).Create(&s).Error
	if err != nil {
		utils.Send([]string{s.App}, "测试报警", "")
	}
	return
}

// 每个level只能对应一个报警
func checkLevelExist(app, level string) bool {
	var count int64
	err := global.GVA_DB.Table(model.GetLevelAlarmTableName(app)).Where("level = ?", level).Count(&count).Error
	return err != nil || count != 0
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteLevelAlarmStrategy
//@description: 删除LevelAlarmStrategy记录
//@param: s model.LevelAlarmStrategy
//@return: err error

func DeleteLevelAlarmStrategy(s model.LevelAlarmStrategy) (err error) {
	err = global.GVA_DB.Table(model.GetLevelAlarmTableName(s.App)).Unscoped().Delete(&s).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteLevelAlarmStrategyByIds
//@description: 批量删除LevelAlarmStrategy记录
//@param: ids request.IdsReq, app string
//@return: err error

func DeleteLevelAlarmStrategyByIds(ids request.IdsReq, app string) (err error) {
	err = global.GVA_DB.Table(model.GetLevelAlarmTableName(app)).Unscoped().Delete(&[]model.LevelAlarmStrategy{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateLevelAlarmStrategy
//@description: 更新LevelAlarmStrategy记录
//@param: s *model.LevelAlarmStrategy
//@return: err error

func UpdateLevelAlarmStrategy(s model.LevelAlarmStrategy) (err error) {
	s.Level = strings.ToLower(s.Level)
	if checkLevelExist(s.App, s.Level) {
		return errors.New("level existed")
	}

	return global.GVA_DB.Table(model.GetLevelAlarmTableName(s.App)).Model(&s).
		Omit("start_time", "start_count").Updates(s).Error
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetLevelAlarmStrategy
//@description: 根据id获取LevelAlarmStrategy记录
//@param: id uint, app string
//@return: err error, s model.LevelAlarmStrategy

func GetLevelAlarmStrategy(id uint, app string) (err error, s model.LevelAlarmStrategy) {
	err = global.GVA_DB.Table(model.GetLevelAlarmTableName(app)).Where("id = ?", id).First(&s).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetLevelAlarmStrategyInfoList
//@description: 分页获取LevelAlarmStrategy记录
//@param: info request.LevelAlarmStrategySearch
//@return: err error, list interface{}, total int64

func GetLevelAlarmStrategyInfoList(info request.LevelAlarmStrategySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(model.GetLevelAlarmTableName(info.App))
	var ss []model.LevelAlarmStrategy
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ss).Error
	return err, ss, total
}
