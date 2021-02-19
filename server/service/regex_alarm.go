package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"regexp"
	"time"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: CreateRegexAlarmStrategy
//@description: 创建RegexAlarmStrategy记录
//@param: s model.RegexAlarmStrategy
//@return: err error

func CreateRegexAlarmStrategy(s model.RegexAlarmStrategy) (err error) {
	if _, found := global.APP_MANAGER.Load(s.App); !found {
		return errors.New("app doesn't exist")
	}
	s.Regexp, err = regexp.Compile(s.Regex)
	if err != nil {
		return
	}
	s.StartTime = time.Now()
	s.StartCount = 0
	err = global.GVA_DB.Table(model.GetRegexAlarmTableName(s.App)).Create(&s).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteRegexAlarmStrategy
//@description: 删除RegexAlarmStrategy记录
//@param: s model.RegexAlarmStrategy
//@return: err error

func DeleteRegexAlarmStrategy(s model.RegexAlarmStrategy) (err error) {
	err = global.GVA_DB.Table(model.GetRegexAlarmTableName(s.App)).Unscoped().Delete(&s).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteRegexAlarmStrategyByIds
//@description: 批量删除RegexAlarmStrategy记录
//@param: ids request.IdsReq, app string
//@return: err error

func DeleteRegexAlarmStrategyByIds(ids request.IdsReq, app string) (err error) {
	err = global.GVA_DB.Table(model.GetRegexAlarmTableName(app)).Unscoped().Delete(&[]model.RegexAlarmStrategy{},"id in ?",ids.Ids).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateRegexAlarmStrategy
//@description: 更新RegexAlarmStrategy记录
//@param: s *model.RegexAlarmStrategy
//@return: err error

func UpdateRegexAlarmStrategy(s model.RegexAlarmStrategy) (err error) {
	s.Regexp, err = regexp.Compile(s.Regex)
	if err != nil {
		return
	}
	err = global.GVA_DB.Table(model.GetRegexAlarmTableName(s.App)).Save(&s).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetRegexAlarmStrategy
//@description: 根据id获取RegexAlarmStrategy记录
//@param: id uint, app string
//@return: err error, s model.RegexAlarmStrategy

func GetRegexAlarmStrategy(id uint, app string) (err error, s model.RegexAlarmStrategy) {
	err = global.GVA_DB.Table(model.GetRegexAlarmTableName(app)).Where("id = ?", id).First(&s).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetRegexAlarmStrategyInfoList
//@description: 分页获取RegexAlarmStrategy记录
//@param: info request.RegexAlarmStrategySearch
//@return: err error, list interface{}, total int64

func GetRegexAlarmStrategyInfoList(info request.RegexAlarmStrategySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Table(model.GetRegexAlarmTableName(info.App))
    var ss []model.RegexAlarmStrategy
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Email != "" {
		db = db.Where("`email` LIKE ?","%"+ info.Email+"%")
    }
    if info.Regex != "" {
		db = db.Where("`regex` LIKE ?","%"+ info.Regex+"%")

	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ss).Error
	return err, ss, total
}