package model

import (
	orm "basic-antd/init/global"
	"errors"
)

// Api 管理
type Api struct {
	ID     int    `gorm:"primary_key;auto_increment" form:"id" json:"id"`
	Path   string `form:"path" json:"path"`
	Group  string `form:"group" json:"group"` // 分组
	Desc   string `form:"desc" json:"desc"`   // 简介
	Method string `form:"method" json:"method"`
	Model
}

func (a *Api) Insert() (err error) {
	var api []Api
	if orm.Eloquent.Where(&Api{Path: a.Path, Method: a.Method}).Find(&api).RowsAffected > 0 {
		return errors.New("路径重复")
	}
	if err = orm.Eloquent.Create(&a).Error; err != nil {
		return
	}
	return
}

func (a *Api) Get() (api Api, err error) {
	if err = orm.Eloquent.First(&api, a.ID).Error; err != nil {
		return
	}
	return
}

func (a *Api) GetPage(pageIndex, pageSize int) (api []Api, count int64, err error) {
	table := orm.Eloquent
	if a.Path != "" {
		table = table.Where("path like ?", "%"+a.Path+"%")
	}

	if a.Desc != "" {
		table = table.Where("desc like ?", "%"+a.Desc+"%")
	}

	if a.Group != "" {
		table = table.Where("group = ?", a.Group)
	}

	if a.Method != "" {
		table = table.Where("method = ?", a.Method)
	}

	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&api).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

func (a *Api) GetList() (api []Api, err error) {
	if err = orm.Eloquent.Find(&api).Error; err != nil {
		return
	}
	return
}

func (a *Api) Update() (err error) {
	if err = orm.Eloquent.Model(&a).Updates(&a).Error; err != nil {
		return
	}
	return
}

func (a *Api) Delete(id []int) (err error) {
	if err = orm.Eloquent.Unscoped().Delete(Api{}, "id in ?", id).Error; err != nil {
		return
	}
	return
}
