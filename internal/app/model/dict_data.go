package model

import (
	orm "basic-antd/init/global"
	"errors"
)

type DictData struct {
	DictCode int    `gorm:"primaryKey;autoIncrement" json:"dictCode"`
	Label    string `gorm:"size:128" json:"label"`
	Value    string `gorm:"size:128" json:"value"`
	Type     string `gorm:"size:64" json:"type"`
	Status   string `gorm:"size:2" json:"status"`
	By
	Model
}

func (DictData) TableName() string {
	return "dict_data"
}

func (d *DictData) Insert() (err error) {
	var count int64
	if err = orm.Eloquent.Table(d.TableName()).
		Where("type = ?", d.Type).
		Where("label = ? or (label = ? and value = ?)", d.Label, d.Label, d.Value).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		return errors.New("字典标签或者字典键值已经存在!")
	}
	if err = orm.Eloquent.Table(d.TableName()).Create(&d).Error; err != nil {
		return
	}
	return
}

func (d *DictData) Get() (dictData DictData, err error) {
	if err = orm.Eloquent.Table(d.TableName()).First(&dictData, d.DictCode).Error; err != nil {
		return
	}
	return
}

func (d *DictData) GetList() (dictData []DictData, count int64, err error) {
	table := orm.Eloquent.Table(d.TableName())
	if d.Type != "" {
		table = table.Where("type = ?", d.Type)
	}
	if err = table.Find(&dictData).Count(&count).Error; err != nil {
		return
	}
	return
}

func (d *DictData) GetPage(pageIndex, pageSize int) (dictData []DictData, count int64, err error) {
	table := orm.Eloquent.Table(d.TableName())
	if d.DictCode != 0 {
		table = table.Where("dict_code = ?", d.DictCode)
	}
	if d.Type != "" {
		table = table.Where("type = ?", d.Type)
	}
	if d.Label != "" {
		table = table.Where("label = ?", d.Label)
	}
	if d.Status != "" {
		table = table.Where("status = ?", d.Status)
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&dictData).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

func (d *DictData) Update() (dictData DictData, err error) {
	if err = orm.Eloquent.Table(d.TableName()).First(&dictData, d.DictCode).Error; err != nil {
		return
	}
	if d.Label != "" && d.Label != dictData.Label {
		return dictData, errors.New("标签不允许修改")
	}
	if d.Value != "" && d.Value != dictData.Value {
		return dictData, errors.New("键值不允许修改")
	}
	if err = orm.Eloquent.Table(d.TableName()).Model(&dictData).Updates(&d).Error; err != nil {
		return
	}
	return
}

func (d *DictData) Delete(id []int) (err error) {
	if err = orm.Eloquent.Table(d.TableName()).Where("dict_code in (?)", id).Delete(&d).Error; err != nil {
		return
	}
	return
}
