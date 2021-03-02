package model

import (
	orm "basic-antd/init/global"
	"basic-antd/tools/encrypt"
	"errors"
	"strconv"
)

type User struct {
	UserId        int           `gorm:"primaryKey;autoIncrement" json:"userId" form:"userId"`
	Username      string        `gorm:"size:64" json:"username" form:"username"` // 用户名
	RealName      string        `json:"realName" form:"realName"`                // 真实姓名
	Avatar        string        `json:"avatar" form:"avatar"`
	Name          string        `gorm:"size:20" json:"name" form:"name"` // 昵称
	Email         string        `gorm:"size:60" json:"email" form:"email"`
	Password      string        `gorm:"size:128" json:"password" form:"password"`
	Phone         string        `gorm:"size:11" json:"phone" form:"phone"`
	LastLoginIp   string        `gorm:"size:64" json:"lastLoginIp" form:"lastLoginIp"`
	LastLoginTime string        `json:"lastLoginTime" form:"lastLoginTime"` // 登录时间戳
	LoginTimes    int           `gorm:"size:225" json:"loginTimes"`
	Sex           string        `gorm:"size:2" json:"sex"`
	Status        string        `json:"status" gorm:"size:2" form:"status"` // 1 启用 2 停用
	Roles         []*Role       `json:"roles" gorm:"-"`
	Departments   []*Department `gorm:"-" json:"departments"`
	RoleIds       []int         `gorm:"-" json:"roleIds"`
	DepartmentIds []int         `gorm:"-" json:"departmentIds"`
	DepartmentId  int           `gorm:"-" json:"departmentId"` // 菜单树 点选条件
	By
	Model
}

type UserInfo struct {
	UserId        int           `gorm:"primary_key;auto_increment" json:"userId"`
	Username      string        `gorm:"size:64" json:"username" form:"username"`
	RealName      string        `json:"realName" form:"realName"`
	Avatar        string        `json:"avatar"`
	Name          string        `gorm:"size:20" json:"name"` // 昵称
	Email         string        `gorm:"size:60" json:"email,omitempty"`
	Phone         string        `gorm:"size:11" json:"phone,omitempty"`
	LastLoginIp   string        `gorm:"size:64" json:"lastLoginIp,omitempty"`
	LastLoginTime string        `json:"lastLoginTime,omitempty"` // 登录时间戳
	LoginTimes    int           `gorm:"size:225" json:"loginTimes"`
	Sex           string        `gorm:"size:2" json:"sex"`
	Status        string        `json:"status" gorm:"size:2"` // 1 启用 2 停用
	Role          []interface{} `json:"role" gorm:"-"`
	Model
}

func (User) TableName() string {
	return "user"
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (l *Login) GetUser() (user User, flag int, err error) {
	if err = orm.Eloquent.Table(user.TableName()).Where("username = ?", l.Username).First(&user).Error; err != nil {
		flag = 1
		return
	}

	if user.Status != "1" {
		flag = 2
		return
	}

	if encrypt.Md5(l.Password) != user.Password {
		flag = 3
		return
	}
	if err = orm.Eloquent.Table(user.TableName()).Select("user_role.*,role.*").Where("user.user_id = ? and role.status = ?", user.UserId, 1).Joins("left join user_role on user.user_id = user_role.user_id").Joins("left join role on role.role_id = user_role.role_id").Order("sort").Find(&user.Roles).Error; err != nil {
		return
	}
	flag = 0
	return
}

func (u *User) GetUserById(userId int) (user User, err error) {
	if err = orm.Eloquent.Table(u.TableName()).Where("user_id = ?", userId).First(&user).Error; err != nil {
		return
	}
	var (
		roles       []Role
		departments []Department
	)
	if err = orm.Eloquent.Select("user_role.*,role.*").Joins("left join user_role on user_role.role_id = role.role_id").Where("user_role.user_id = ?", userId).Find(&roles).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Select("department_user.*, department.name").Joins("left join department_user on department_user.department_id = department.department_id").Where("department_user.user_id = ?", userId).Find(&departments).Error; err != nil {
		return
	}
	for _, v := range roles {
		user.RoleIds = append(user.RoleIds, v.RoleId)
	}
	for _, v := range departments {
		user.DepartmentIds = append(user.DepartmentIds, v.DepartmentId)
	}
	return
}

func (u *UserInfo) GetUserInfo(userId int) (user UserInfo, err error) {
	if err = orm.Eloquent.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserInfo) GetUserAndRoleById() (user UserInfo, err error, msg string) {
	if err = orm.Eloquent.Table("user").Where("user_id = ?", u.UserId).First(&user).Error; err != nil {
		return user, err, "查询用户id失败"
	}
	return
}

func (u *User) GetPage(pageSize, pageIndex int) (users []User, count int64, err error) {
	var (
		du      []DepartmentUser
		userIds []int
	)
	table := orm.Eloquent.Table("user").Where("`deleted_at` IS NULL")
	if u.Name != "" {
		table = table.Where("name LIKE ?", "%"+u.Name+"%")
	}
	if u.Username != "" {
		table = table.Where("username LIKE ?", "%"+u.Username+"%")
	}
	if u.Phone != "" {
		table = table.Where("phone = ?", u.Phone)
	}
	if u.DepartmentId > 0 {
		if err = orm.Eloquent.Table("department_user").Where("department_id = ?", u.DepartmentId).Find(&du).Error; err != nil {
			return
		}
		for _, v := range du {
			userIds = append(userIds, v.UserId)
		}
		table = table.Where("user_id in (?)", userIds)
	}
	rows, err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Offset(-1).Limit(-1).Count(&count).Rows()
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var user User
		err = orm.Eloquent.ScanRows(rows, &user)
		if err != nil {
			return
		}
		if err = orm.Eloquent.Select("user_role.*,role.*").Joins("left join user_role on user_role.role_id = role.role_id").Where("user_role.user_id = ?", user.UserId).Find(&user.Roles).Error; err != nil {
			return
		}
		if err = orm.Eloquent.Select("department_user.*, department.name").Joins("left join department_user on department_user.department_id = department.department_id").Where("department_user.user_id = ?", user.UserId).Find(&user.Departments).Error; err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (u *User) Insert() (userId int, err error) {
	var count int64
	table := orm.Eloquent.Table(u.TableName())
	if err = table.Where("username = ?", u.Username).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		err = errors.New("用户已存在")
		return
	}
	if err = table.Create(&u).Error; err != nil {
		return
	}
	return u.UserId, err
}

func (u *User) Update() (user User, err error) {
	if err = orm.Eloquent.Table(u.TableName()).First(&user, u.UserId).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Table(u.TableName()).Model(&user).Updates(&u).Error; err != nil {
		return
	}
	return
}

func (u *User) BatchDelete(id []int) (err error) {
	if err = orm.Eloquent.Table(u.TableName()).Where("user_id in (?)", id).Delete(&User{}).Error; err != nil {
		return
	}
	return
}

/* 获取user的 角色 和 角色 id 以 逗号隔开*/
func (u *User) GetRoleStr() (roles, roleIds string) {
	if len(u.Roles) > 0 {
		for i := 0; i < len(u.Roles); i++ {
			if i == len(u.Roles)-1 {
				roles += u.Roles[i].RoleKey
				roleIds += strconv.Itoa(u.Roles[i].RoleId)
			} else {
				roles += u.Roles[i].RoleKey + ","
				roleIds += strconv.Itoa(u.Roles[i].RoleId) + ","
			}
		}
	}
	return
}
