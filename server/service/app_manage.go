package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [Durden-T](https://github.com/Durden-T.)
//@function: CreateApp
//@description: 创建app
//@param: e model.App
//@return: err error

func CreateApp(e *model.App) (err error) {
	_, ok := global.APP_MANAGER.Load(e.Name)
	if ok {
		return errors.New("app exist")
	}

	if err := e.Init(); err != nil {
		return err
	}

	err = global.GVA_DB.Create(e).Error
	if err != nil {
		return
	}

	global.APP_MANAGER.Store(e.Name, e)
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: DeleteApp
//@description: 删除app
//@param: e model.App
//@return: err error

func DeleteApp(e *model.App) (err error) {
	appInterface, ok := global.APP_MANAGER.Load(e.Name)
	if !ok {
		return errors.New("app not exist")
	}
	app, ok := appInterface.(*model.App)
	if !ok {
		return errors.New("app not exist")
	}

	app.Stop()
	err = global.GVA_DB.Delete(e).Error
	if err != nil {
		return
	}
	global.APP_MANAGER.Delete(e.Name)
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetAppInfoList
//@description: 分页获取app列表
//@param: info request.PageInfo
//@return: list []model.App, total int64, err error

func GetAppInfoList(info *request.PageInfo) (list []model.App, total int64, err error) {
	db := global.GVA_DB.Model(&model.App{})
	err = db.Count(&total).Error
	if err != nil || total == 0 {
		return
	}

	if info != nil {
		limit := info.PageSize
		offset := info.PageSize * (info.Page - 1)
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error
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

	var dest model.SysAuthority
	err = global.GVA_DB.Model(&model.SysAuthority{}).Where("authority_id = ?", sysUserAuthorityID).
		Preload("App", "name = ?", appName).Take(&dest).Error
	if err == nil && len(dest.App) > 0 {
		pass = true
	}
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: GetAllApps
//@description: 获取所有app信息
//@param:
//@return: appList []model.App, err error

func GetAllApps() (appList []*model.App, err error) {
	err = global.GVA_DB.Model(&model.App{}).Find(&appList).Error
	return
}

//@author: [Durden-T](https://github.com/Durden-T)
//@function: UpdateApp
//@description: 更新app, 是否启用报警
//@param: app *model.App
//@return: appList []model.App, err error

func UpdateApp(app *model.App) error {
	appInterface, ok := global.APP_MANAGER.Load(app.Name)
	if !ok {
		return errors.New("app not exist")
	}
	oldApp, ok := appInterface.(*model.App)
	if !ok {
		return errors.New("app not exist")
	}

	if app.EnableAlarm == oldApp.EnableAlarm {
		return nil
	}

	if app.EnableAlarm {
		if err := oldApp.InitAlarm(); err != nil {
			return err
		}
	} else {
		oldApp.DisableAlarm()
	}

	return global.GVA_DB.Model(&model.App{}).Where("name = ?", app.Name).
		Update("enable_alarm", app.EnableAlarm).Error
}
