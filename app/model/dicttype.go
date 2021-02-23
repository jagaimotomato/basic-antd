package model

import (
	orm "basic-antd/init/global"
	"errors"
)

type DictType struct {
	DictId int    `gorm:"primaryKey;autoIncrement" json:"dictId"`
	Name   string `gorm:"size:128" json:"name"`
	Type   string `gorm:"size:128" json:"type"`
	Status string `gorm:"size:2" json:"status"`
	By
	Model
}

func (DictType) TableName() string {
	return "dict_type"
}

func (d *DictType) Insert() (err error) {
	var count int64
	if err = orm.Eloquent.Table(d.TableName()).Where("name = ? or type = ?", d.Name, d.Type).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		return errors.New("字典名称或者字典类型已经存在")
	}
	if err = orm.Eloquent.Table(d.TableName()).Create(&d).Error; err != nil {
		return
	}
	return
}

func (d *DictType) Get() (dictType DictType, err error) {
	if err = orm.Eloquent.Table(d.TableName()).First(&dictType, d.DictId).Error; err != nil {
		return
	}
	return
}

func (d *DictType) GetList() (dictTypes []DictType, count int64, err error) {
	if err = orm.Eloquent.Table(d.TableName()).Find(&dictTypes).Count(&count).Error; err != nil {
		return
	}
	return
}

func (d *DictType) GetPage(pageIndex, pageSize int) (dictTypes []DictType, count int64, err error) {
	table := orm.Eloquent.Table(d.TableName())
	if d.DictId != 0 {
		table = table.Where("dict_id = ?", d.DictId)
	}
	if d.Name != "" {
		table = table.Where("name = ?", d.Name)
	}
	if d.Type != "" {
		table = table.Where("type = ?", d.Type)
	}
	if d.Status != "" {
		table = table.Where("status = ?", d.Status)
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&dictTypes).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

func (d *DictType) Update() (dictType DictType, err error) {
	if err = orm.Eloquent.Table(d.TableName()).First(&dictType, d.DictId).Error; err != nil {
		return
	}
	if d.Name != "" && d.Name != dictType.Name {
		return dictType, errors.New("名称不允许修改")
	}

	if d.Type != "" && d.Type != dictType.Type {
		return dictType, errors.New("类型不允许修改")
	}

	if err = orm.Eloquent.Table(d.TableName()).Model(&dictType).Updates(&d).Error; err != nil {
		return
	}
	return
}

func (d *DictType) Delete(id []int) (err error) {
	if err = orm.Eloquent.Table(d.TableName()).Where("dict_id in (?)", id).Delete(&d).Error; err != nil {
		return
	}
	return
}
