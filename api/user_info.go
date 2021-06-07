package api

import (
	"basic-antd/internal/app/model"
	"basic-antd/pkg/app"
	"basic-antd/pkg/app/msg"
	"basic-antd/pkg/jwt"
	"basic-antd/tools"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

func GetUserInfo(c *gin.Context) {
	var (
		userInfo model.UserInfo
		role     model.Role
		p        model.Permission
	)
	userId := jwt.GetUserId(c)
	fmt.Println(userId, "***************")
	roleIds := tools.StringGroupToIntGroup(strings.Split(jwt.GetRoleIdStr(c), ","))
	/* todo 登录时将该user存入redis 获取信息直接取出*/
	userInfo, err := userInfo.GetUserInfo(userId)
	tools.HasError(err, msg.Failed, 500)
	roles, err := role.GetRolesByIds(roleIds)
	tools.HasError(err, msg.Failed, 500)
	/* 组装数据 */
	for _, v := range roles {
		role := make(map[string]interface{})
		permissions, err := p.GetPermissionsByRoleId(v.RoleId)
		tools.HasError(err, msg.Failed, 500)
		p := make([]interface{}, 0)
		for i := 0; i < len(permissions); i++ {
			if permissions[i].Actions == "" {
				permission := make(map[string]interface{})
				permission["roleId"] = v.RoleKey
				permission["permissionId"] = permissions[i].Name
				permission["permissionName"] = permissions[i].Title
				actionEntitySet := make([]interface{}, 0)
				actionList := make([]string, 0)
				for j := 0; j < len(permissions); j++ {
					if permissions[j].Actions != "" {
						if permissions[i].PermissionId == permissions[j].ParentId {
							actionEntitySetMap := make(map[string]interface{})
							actionEntitySetMap["defaultCheck"] = false
							actionEntitySetMap["describe"] = permissions[j].Title
							actionEntitySetMap["action"] = permissions[j].Name
							actionEntitySet = append(actionEntitySet, actionEntitySetMap)
							actionList = append(actionList, permissions[j].Name)
						}
					}
				}
				if len(actionEntitySet) > 0 {
					permission["actionEntitySet"] = actionEntitySet
					permission["actionList"] = actionList
				}
				p = append(p, permission)
			}
		}
		role["permissions"] = p
		role["status"] = v.Status
		role["describe"] = v.Describe
		role["name"] = v.Name
		role["id"] = v.RoleKey
		userInfo.Role = append(userInfo.Role, role)
	}
	app.Success(c, userInfo, msg.Success)
}

func GetUserNav(c *gin.Context) {
	var (
		menu    model.Permission
		menus   []model.Permission
		isAdmin = false
		err     error
	)
	for _, v := range strings.Split(jwt.GetRoleName(c), ",") {
		if v == "admin" {
			isAdmin = true
			break
		}
	}
	if isAdmin {
		menus, err = menu.GetAdminMenus()
		tools.HasError(err, msg.Failed, 500)
	} else {
		roleIds := tools.StringGroupToIntGroup(strings.Split(jwt.GetRoleIdStr(c), ","))
		menus, err = menu.GetMenus(roleIds)
		tools.HasError(err, msg.Failed, -1)
	}
	app.Success(c, menus, msg.Success)
}

/* 个人中心获取本人信息 */
func GetAccountInfo(c *gin.Context) {
	var (
		u   model.User
		err error
	)
	user, err := u.GetUserById(jwt.GetUserId(c))
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, user, msg.Success)
}

/* 更新密码 */
func UpdateAccountPassword(c *gin.Context) {
	var (
		u   model.User
		p   model.Password
		err error
	)
	err = c.MustBindWith(&p, binding.JSON)
	tools.HasError(err, msg.BindingFailed, 400)
	user, err := u.GetUserById(jwt.GetUserId(c))
	tools.HasError(err, msg.Failed, 500)
	if !tools.Verify(user.Password, p.Old, user.Salt) {
		app.ResponseError(c, 403, errors.New("旧密码错误"), "")
		return
	}
	u.Password, u.Salt, err = tools.GeneratePassword(p.New)
	u.PasswordStrength = p.PasswordStrength
	u.UserId = jwt.GetUserId(c)
	user, err = u.Update()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, user, msg.Success)
}

/* 更新个人信息 */
func UpdateAccountInfo(c *gin.Context) {
	var (
		u   model.User
		err error
	)
	err = c.MustBindWith(&u, binding.JSON)
	tools.HasError(err, msg.BindingFailed, -1)
	u.UserId = jwt.GetUserId(c)
	user, err := u.Update()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, user, msg.Success)
}

func GetSaltPassword(c *gin.Context) {
	pass, salt, _ := tools.GeneratePassword("123456")
	app.Custom(c, gin.H{
		"password": pass,
		"salt":     salt,
	})
}
