package model

import (
	orm "basic-antd/init/global"
	"basic-antd/tools/system"
)

type RolePermission struct {
	RoleId       int    `json:"roleId"`
	PermissionId int    `json:"permissionId"`
	RoleKey      string `json:"roleKey"`
	By
}

func (*RolePermission) TableName() string {
	return "role_permission"
}

func GetRolePermissionIds(roleId int) (permissionIds []int, err error) {
	var (
		permission []Permission
	)
	if err = orm.Eloquent.Select("role_permission.*, permission.permission_id, permission.actions, permission.parent_id").Where("role_id = ?", roleId).Joins("left join role_permission on permission.permission_id = role_permission.permission_id").Find(&permission).Error; err != nil {
		return
	}
	var arr []system.Child
	for _, v := range permission {
		var child system.Child
		child.Id = v.PermissionId
		child.Pid = v.ParentId
		arr = append(arr, child)
	}
	/* 如果有子菜单就push子菜单，父级不push*/
	for _, v := range permission {
		if system.HasChild(arr, v.PermissionId) {
			continue
		}
		permissionIds = append(permissionIds, v.PermissionId)
	}
	return
}

func (rp *RolePermission) Insert(roleId int, permissionId []int) (err error) {
	var (
		role        Role
		permissions []Permission
		rps         = make([]RolePermission, 0)
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

	if err = tx.Table(role.TableName()).Where("role_id = ?", roleId).First(&role).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Table("permission").Where("permission_id in (?)", permissionId).Find(&permissions).Error; err != nil {
		tx.Rollback()
		return
	}
	// rolePermissionSql := "INSERT INTO `role_permission` (`role_id`,`permission_id`,`role_key`) VALUES "
	for _, v := range permissions {
		var rp RolePermission
		rp.RoleKey = role.RoleKey
		rp.PermissionId = v.PermissionId
		rp.RoleId = role.RoleId
		rps = append(rps, rp)
	}
	if err = orm.Eloquent.Table(rp.TableName()).Create(&rps).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

func (rp *RolePermission) DeleteRolePermission(roleId []int) (err error) {
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return
	}
	if err = tx.Table(rp.TableName()).Where("role_id in (?)", roleId).Delete(&rp).Error; err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}
