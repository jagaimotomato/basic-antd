package model

import (
	orm "basic-antd/init/global"
	"errors"
)

type Role struct {
	RoleId        int          `gorm:"primaryKey;autoIncrement" json:"roleId"`
	RoleKey       string       `gorm:"size:255" json:"roleKey"` // "admin"
	Name          string       `gorm:"size:255" json:"name"`    // 管理员
	Describe      string       `json:"describe"`                // 描述
	Status        string       `gorm:"size:2" json:"status"`
	Sort          int          `json:"sort"`
	PermissionIds []int        `gorm:"-" json:"permissionIds"`
	Apis          []CasbinInfo `gorm:"-" json:"apis"`
	Permission    interface{}  `gorm:"-" json:"permissions"`
	By
	Model
}

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type RoleInfo struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Describe    string        `json:"describe"`
	Status      string        `json:"status"`
	Permissions []interface{} `gorm:"-" json:"permissions"`
}

func (*Role) TableName() string {
	return "role"
}

func (r *Role) GetRolesByIds(roleIds []int) (roles []Role, err error) {
	if err = orm.Eloquent.Table(r.TableName()).Where("role_id in (?)", roleIds).Find(&roles).Error; err != nil {
		return
	}
	return
}

func (r *Role) GetList() (role []Role, err error) {
	if err = orm.Eloquent.Table(r.TableName()).Find(&role).Error; err != nil {
		return
	}
	return
}

func (r *Role) GetPage(pageIndex, pageSize int) (role []Role, count int64, err error) {
	table := orm.Eloquent.Table(r.TableName())
	if r.RoleKey != "" {
		table = table.Where("role_key = ?", r.RoleKey)
	}
	if r.Name != "" {
		table = table.Where("name = ?", r.Name)
	}
	if r.Describe != "" {
		table = table.Where("describe like ?", "%"+r.Describe+"%")
	}
	if r.Status != "" {
		table = table.Where("status = ?", r.Status)
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("sort").Find(&role).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return
	}
	return
}

func (r *Role) GetRolePermissionPage(pageIndex, pageSize int) (rolePage []Role, count int64, err error) {
	rolePage, count, err = r.GetPage(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(rolePage) > 0 {
		for k, v := range rolePage {
			var (
				permissions    []Permission
				newPermissions = make([]Permission, 0)
			)
			if err = orm.Eloquent.Select("role_permission.*, permission.*").Where("role_id = ?", v.RoleId).Joins("left join role_permission on permission.permission_id = role_permission.permission_id").Find(&permissions).Order("sort").Error; err != nil {
				return
			}
			for _, v := range permissions {
				if v.Actions != "" {
					continue
				}
				newPermissions = append(newPermissions, v)
			}
			for i := 0; i < len(newPermissions); i++ {
				for j := 0; j < len(permissions); j++ {
					if newPermissions[i].PermissionId != permissions[j].ParentId {
						continue
					}
					if permissions[j].Actions != "" {
						newPermissions[i].ActionList = append(newPermissions[i].ActionList, permissions[j].Actions)
					}
				}
			}
			rolePage[k].Permission = newPermissions
		}
	}
	return
}

func (r *Role) Get() (role Role, err error) {
	if err = orm.Eloquent.Table(r.TableName()).First(&role, r.RoleId).Error; err != nil {
		return
	}
	return
}

func (r *Role) Insert() (roleId int, err error) {
	var i int64
	orm.Eloquent.Table(r.TableName()).Where("role_name = ? or role_key = ?", r.Name, r.RoleKey).Count(&i)
	if i > 0 {
		return 0, errors.New("角色名称或者角色key已经存在")
	}
	if err = orm.Eloquent.Table(r.TableName()).Create(&r).Error; err != nil {
		return
	}
	return r.RoleId, err
}

func (r *Role) Update() (err error) {
	if err = orm.Eloquent.Table(r.TableName()).Model(&r).Updates(&r).Error; err != nil {
		return
	}
	return
}

func (r *Role) Delete(roleId []int) (err error) {
	tx := orm.Eloquent.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var count int64
	if err := tx.Table("user_role").Where("role_id in (?)", roleId).Count(&count).Error; err != nil {
		tx.Rollback()
		return err
	}

	if count > 0 {
		tx.Rollback()
		return errors.New("存在绑定用户，请先解绑后再删除")
	}

	if err := tx.Table(r.TableName()).Where("role_id in (?)", roleId).Unscoped().Delete(&Role{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
