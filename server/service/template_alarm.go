package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"time"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: CreateTemplateAlarmStrategy
//@description: 创建TemplateAlarmStrategy记录
//@param: s model.TemplateAlarmStrategy
//@return: err error

func CreateTemplateAlarmStrategy(s model.TemplateAlarmStrategy) (err error) {
	if _, found := global.APP_MANAGER.Load(s.App); !found {
		return errors.New("app doesn't exist")
	}
	var template model.LogTemplate
	if err = global.GVA_DB.Table(model.GetTemplateTableName(s.App)).Where("cluster_id = ?", s.TemplateId).Take(&template).Error; err != nil {
		return
	}
	now := time.Now()
	s.StartTime = now
	s.StartCount = template.Size

	err = global.GVA_DB.Table(model.GetTemplateAlarmTableName(s.App)).Create(&s).Error
	if err != nil {
		return
	}
	return utils.Send([]string{s.Email}, "测试报警", "")
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteTemplateAlarmStrategy
//@description: 删除TemplateAlarmStrategy记录
//@param: s model.TemplateAlarmStrategy
//@return: err error

func DeleteTemplateAlarmStrategy(s model.TemplateAlarmStrategy) (err error) {
	err = global.GVA_DB.Table(model.GetTemplateAlarmTableName(s.App)).Delete(&s).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteTemplateAlarmStrategyByIds
//@description: 批量删除TemplateAlarmStrategy记录
//@param: ids request.IdsReq app string
//@return: err error

func DeleteTemplateAlarmStrategyByIds(ids request.IdsReq, app string) (err error) {
	err = global.GVA_DB.Table(model.GetTemplateAlarmTableName(app)).Delete(&[]model.TemplateAlarmStrategy{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateTemplateAlarmStrategy
//@description: 更新TemplateAlarmStrategy记录
//@param: s *model.TemplateAlarmStrategy
//@return: err error

func UpdateTemplateAlarmStrategy(s model.TemplateAlarmStrategy) (err error) {
	return global.GVA_DB.Table(model.GetTemplateAlarmTableName(s.App)).Model(&s).
		Omit("start_time", "start_count").Updates(&s).Error
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTemplateAlarmStrategy
//@description: 根据id获取TemplateAlarmStrategy记录
//@param: id uint, app string
//@return: err error, s model.TemplateAlarmStrategy

func GetTemplateAlarmStrategy(id uint, app string) (err error, s model.TemplateAlarmStrategy) {
	err = global.GVA_DB.Table(model.GetTemplateAlarmTableName(app)).Where("id = ?", id).First(&s).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetTemplateAlarmStrategyInfoList
//@description: 分页获取TemplateAlarmStrategy记录
//@param: info request.TemplateAlarmStrategySearch
//@return: err error, list interface{}, total int64

func GetTemplateAlarmStrategyInfoList(info request.TemplateAlarmStrategySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(model.GetTemplateAlarmTableName(info.App))
	var ss []model.TemplateAlarmStrategy
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.TemplateId != 0 {
		db = db.Where("`template_id` = ?", info.TemplateId)
	}
	if info.Email != "" {
		db = db.Where("`email` LIKE ?", "%"+info.Email+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&ss).Error
	return err, ss, total
}
