package api

import (
	"basic-antd/app/model"
	"basic-antd/pkg/jwt"
	"basic-antd/tools"
	"basic-antd/tools/app"
	"basic-antd/tools/app/msg"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetUserInfo(c *gin.Context) {
	var (
		userInfo model.UserInfo
		role     model.Role
		p        model.Permission
	)
	userId := jwt.GetUserId(c)
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
							actionEntitySetMap["action"] = permissions[j].Actions
							actionEntitySet = append(actionEntitySet, actionEntitySetMap)
							actionList = append(actionList, permissions[j].Actions)
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
		menu model.Permission
	)
	roleIds := tools.StringGroupToIntGroup(strings.Split(jwt.GetRoleIdStr(c), ","))
	menus, err := menu.GetMenus(roleIds)
	tools.HasError(err, msg.Failed, -1)
	app.Success(c, menus, msg.Success)
}
