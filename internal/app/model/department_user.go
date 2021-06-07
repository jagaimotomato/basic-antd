package model

import orm "basic-antd/init/global"

type DepartmentUser struct {
	UserId       int
	DepartmentId int
}

func (*DepartmentUser) TableName() string {
	return "department_user"
}

func (du *DepartmentUser) Insert(userId int, departmentIds []int) (err error) {
	var (
		user        User
		departments []Department
		dus         = make([]DepartmentUser, 0)
	)
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return
	}

	if err = tx.Table(user.TableName()).Where("user_id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Table("department").Where("department_id in (?)", departmentIds).Find(&departments).Error; err != nil {
		tx.Rollback()
		return
	}

	for _, v := range departments {
		var du DepartmentUser
		du.UserId = user.UserId
		du.DepartmentId = v.DepartmentId
		dus = append(dus, du)
	}
	if err = orm.Eloquent.Table(du.TableName()).Create(&dus).Error; err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func (du *DepartmentUser) DeleteDepartmentUser(userId []int) (err error) {
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return
	}
	if err = tx.Table(du.TableName()).Where("user_id in (?)", userId).Delete(&du).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
