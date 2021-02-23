package model

import orm "basic-antd/init/global"

type UserRole struct {
	UserId  int    `json:"userId"`
	RoleId  int    `json:"roleId"`
	RoleKey string `json:"roleKey"`
	By
}

func (*UserRole) TableName() string {
	return "user_role"
}

func (ur *UserRole) Insert(userId int, roleIds []int) (err error) {
	var (
		user  User
		roles []Role
		urs   = make([]UserRole, 0)
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

	if err = tx.Table("role").Where("role_id in (?)", roleIds).Find(&roles).Error; err != nil {
		tx.Rollback()
		return
	}

	for _, v := range roles {
		var ur UserRole
		ur.UserId = user.UserId
		ur.RoleId = v.RoleId
		ur.RoleKey = v.RoleKey
		urs = append(urs, ur)
	}

	if err = orm.Eloquent.Table(ur.TableName()).Create(&urs).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func (ur *UserRole) DeleteUserRole(userId []int) (err error) {
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Table(ur.TableName()).Where("user_id in (?)", userId).Delete(&ur).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
