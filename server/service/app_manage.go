package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

//@author: [Durden-T](https://github.com/Durden-T.)
//@function: CreateApp
//@description: 创建app
//@param: e model.App
//@return: err error

func CreateApp(e *model.App) (err error) {
	err = global.GVA_DB.Create(e).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteFileChunk
//@description: 删除app
//@param: e model.App
//@return: err error

func DeleteApp(e *model.App) (err error) {
	err = global.GVA_DB.Delete(e).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateApp
//@description: 更新app
//@param: e *model.App
//@return: err error

func UpdateApp(e *model.App) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetApp
//@description: 获取app信息
//@param: id uint
//@return: err error, app model.App

func GetApp(id uint) (err error, app *model.App) {
	err = global.GVA_DB.Where("id = ?", id).Take(app).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetAppInfoList
//@description: 分页获取app列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func GetAppInfoList(sysUserAuthorityID string, info *request.PageInfo) (list interface{}, total int64, err error) {
	db := global.GVA_DB.Model(&model.App{})
	var a model.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	err, auth := GetAuthorityInfo(a)
	if err != nil {
		return
	}

	var dataId []string
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var appList []model.App
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return
	}
	if info != nil {
		limit := info.PageSize
		offset := info.PageSize * (info.Page - 1)
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&appList).Error

	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: CheckAppPermission
//@description: 检查角色是否有某app的权限
//@param: sysUserAuthorityID string, appName string
//@return: pass bool, err error

func CheckAppPermission(sysUserAuthorityID, appName string) (pass bool, err error) {
	if len(sysUserAuthorityID) == 0 || len(appName) == 0 {
		return false, nil
	}
	err = global.GVA_DB.Where(&model.App{
		SysUserAuthorityID: sysUserAuthorityID,
		Name:               appName,
	}).Take(nil).Error
	if err == nil {
		pass = true
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		pass = false
		err = nil
	}
	return
}
