package model

import (
	orm "basic-antd/init/global"
	"basic-antd/tools/system"
	"errors"
)

type Department struct {
	DepartmentId int          `gorm:"primaryKey;autoIncrement" json:"departmentId"`
	Name         string       `gorm:"size:100" json:"name"`
	ParentId     int          `gorm:"type:int" json:"parentId"`
	Icon         string       `gorm:"size:100" json:"icon"`
	Describe     string       `json:"describe"`
	Children     []Department `gorm:"-" json:"children"`
	By
	Model
}

func (*Department) TableName() string {
	return "department"
}

func (d *Department) Insert() (err error) {
	var departments []Department
	if count := orm.Eloquent.Table(d.TableName()).Where("parent_id = ? and name = ?", 0, d.Name).Find(&departments).RowsAffected; count > 0 {
		return errors.New("存在同名部门")
	}
	if err = orm.Eloquent.Table(d.TableName()).Create(&d).Error; err != nil {
		return
	}
	return
}

func (d *Department) Update() (err error) {
	if err = orm.Eloquent.Table(d.TableName()).Model(&d).Updates(&d).Error; err != nil {
		return
	}
	return
}

func (d *Department) Get() (department Department, err error) {
	if err = orm.Eloquent.Table(d.TableName()).First(&department, d.DepartmentId).Error; err != nil {
		return
	}
	return
}

func (d *Department) GetList() (departments []Department, err error) {
	table := orm.Eloquent.Table(d.TableName())
	if d.Name != "" {
		table = table.Where("name like ?", "%"+d.Name+"%")
	}
	if err = table.Find(&departments).Error; err != nil {
		return
	}
	return
}

func (d *Department) GetPage(pageIndex, pageSize int) (departments []Department, count int64, err error) {
	var (
		list     []Department
		pageList []Department
		arr      []system.Child
	)
	if list, err = d.GetList(); err != nil {
		return
	}
	if err = orm.Eloquent.Table(d.TableName()).Where("parent_id = 0").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&pageList).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	for _, v := range list {
		var child system.Child
		child.Id = v.DepartmentId
		child.Pid = v.ParentId
		arr = append(arr, child)
	}
	for k, _ := range pageList {
		children := Recursion(&list, &pageList[k], &arr)
		departments = append(departments, children)
	}
	return
}

func (d *Department) GetDepartmentTree() (departments []Department, err error) {
	list, err := d.GetList()
	var arr []system.Child
	for _, v := range list {
		var child system.Child
		child.Id = v.DepartmentId
		child.Pid = v.ParentId
		arr = append(arr, child)
	}
	for k, v := range list {
		if v.ParentId != 0 {
			continue
		}
		children := Recursion(&list, &list[k], &arr)
		departments = append(departments, children)
	}
	return
}

func Recursion(departments *[]Department, department *Department, arr *[]system.Child) Department {
	list := *departments
	children := make([]Department, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != department.DepartmentId {
			continue
		}
		m := list[i]
		if system.HasChild(*arr, m.DepartmentId) {
			ms := Recursion(departments, &m, arr)
			children = append(children, ms)
		} else {
			children = append(children, m)
		}
	}
	if len(children) > 0 {
		department.Children = children
	} else {
		department.Children = nil
	}
	return *department
}

/*func HasChild(departments *[]Department, department *Department) bool {
	list := *departments
	for _, v := range list {
		if v.ParentId == department.DepartmentId {
			return true
		}
	}
	return false
}*/

func (d *Department) Delete(id []int) error {
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	// 物理删除
	if err := tx.Table(d.TableName()).Where("department_id in (?)", id).Unscoped().Delete(&Department{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 物理删除
	if err := tx.Table(d.TableName()).Where("parent_id in (?)", id).Unscoped().Delete(&Department{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table("department_user").Where("department_id in (?)", id).Delete(&DepartmentUser{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
