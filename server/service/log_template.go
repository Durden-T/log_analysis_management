package service

import (
	"bytes"
	"errors"

	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	jsoniter "github.com/json-iterator/go"
)

//@author: [Durden-T](https://github.com/Durden-T)
//@function: CreateLogTemplate
//@description: 创建LogTemplate记录
//@param: l model.LogTemplate
//@return: err error

func CreateLogTemplate(l model.LogTemplate) (err error) {
	appInterface, ok := global.APP_MANAGER.Load(l.App)
	if !ok {
		return errors.New("illegal app")
	}
	app := appInterface.(*model.App)
	if app == nil {
		return errors.New("illegal app")
	}

	if !app.EnableAlarm {
		return errors.New("未启用报警")
	}
	buf := bytes.NewBuffer([]byte{})
	enc := jsoniter.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err = enc.Encode(l); err != nil {
		return
	}
	return app.TemplateCollector.ProcessLog(buf.Bytes())

}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteLogTemplate
//@description: 删除LogTemplate记录
//@param: l model.LogTemplate
//@return: err error

func DeleteLogTemplate(l model.LogTemplate) (err error) {
	err = global.GVA_DB.Table(model.GetTemplateTableName(l.App)).Delete(&l).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteLogTemplateByIds
//@description: 批量删除LogTemplate记录
//@param: ids request.IdsReq, app string
//@return: err error

func DeleteLogTemplateByIds(ids request.IdsReq, app string) (err error) {
	err = global.GVA_DB.Table(model.GetTemplateTableName(app)).Delete(&[]model.LogTemplate{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetLogTemplateInfoList
//@description: 分页获取LogTemplate记录
//@param: info request.LogTemplateSearch
//@return: err error, list interface{}, total int64

func GetLogTemplateInfoList(info request.LogTemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(model.GetTemplateTableName(info.App))
	var ls []model.LogTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ClusterId != 0 {
		db = db.Where("`cluster_id` = ?", info.ClusterId)
	}
	if info.Tokens != "" {
		db = db.Where("`tokens` LIKE ?", "%"+info.Tokens+"%")
	}
	if info.Level != "" {
		db = db.Where("`level` = ?", info.Level)
	}
	if info.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+info.Content+"%")
	}
	if info.Tag != "" {
		db = db.Where("`tag` LIKE ?", "%"+info.Tag+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ls).Error
	list = ls
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateLogTemplate
//@description: 更新日志模板
//@param: l model.LogTemplate
//@return: err error

func UpdateLogTemplate(l model.LogTemplate) (err error) {
	return global.GVA_DB.Table(model.GetTemplateTableName(l.App)).Model(&l).
		UpdateColumn("tag", l.Tag).Error
}

