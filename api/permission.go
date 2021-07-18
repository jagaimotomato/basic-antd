package api

import (
	"basic-antd/model"
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/response"
	"basic-antd/pkg/response/msg"
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/* 分页列表 */
func GetPermissionPage(c *gin.Context) {
	var (
		p         model.Permission
		pageIndex = 1
		pageSize  = 10
	)
	index := c.Request.FormValue("page_index")
	if index != "" {
		pageIndex, _ = tools.StringToInt(index)
	}
	size := c.Request.FormValue("page_size")
	if size != "" {
		pageSize, _ = tools.StringToInt(size)
	}
	p.Title = c.Request.FormValue("title")
	permissions, count, err := p.GetRecursionPermission(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	response.Paginate(c, permissions, count, pageIndex, pageSize, msg.Success)
}

/* 选择框数据 */
func GetPermissionList(c *gin.Context) {
	var (
		p model.Permission
	)
	p.Title = c.Request.FormValue("title")
	permissions, err := p.SetPermission()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, permissions, msg.Success)
}

func GetPermission(c *gin.Context) {
	var (
		p model.Permission
	)
	Id, err := tools.StringToInt(c.Param("permissionId"))
	tools.HasError(err, msg.Failed, -1)
	p.PermissionId = Id
	permission, err := p.GetPermission()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, permission, msg.Success)
}

func GetPermissionTree(c *gin.Context) {
	var (
		p model.Permission
	)
	permissions, err := p.GetAllPermission()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, permissions, msg.Success)
}

func GetRolePermissionTree(c *gin.Context) {
	roleId, err := tools.StringToInt(c.Param("roleId"))
	tools.HasError(err, msg.Failed, -1)
	permissionIds, err := model.GetRolePermissionIds(roleId)
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, permissionIds, msg.Success)
}

func CreatePermission(c *gin.Context) {
	var (
		p model.Permission
		// ps []model.Permission
	)
	err := c.MustBindWith(&p, binding.JSON)
	tools.HasError(err, msg.Failed, -1)
	p.CreatedBy = jwt.GetUserIdStr(c)
	_, err = p.Insert()
	/*if len(p.ActionUpload) > 0 {
		for _, v := range p.ActionUpload {
			var permission model.Permission
			permission.Actions = v.Value
			permission.Describe = v.Label
			permission.Title = v.Label
			permission.ParentId = permissionId
			ps = append(ps, permission)
		}
		err = p.BatchInsert(&ps)
	}*/
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}

func UpdatePermission(c *gin.Context) {
	var (
		p model.Permission
		//ps []model.Permission
	)
	err := c.MustBindWith(&p, binding.JSON)
	tools.HasError(err, msg.Failed, -1)
	p.UpdatedBy = jwt.GetUserIdStr(c)
	/*if len(p.ActionUpload) > 0 {
		for _, v := range p.ActionUpload {
			var permission model.Permission
			permission.Actions = v.Value
			permission.Describe = v.Label
			permission.Title = v.Label
			permission.ParentId = p.PermissionId
			ps = append(ps, permission)
		}
		err = p.DeleteActionByPid()
		tools.HasError(err, msg.Failed, 500)
		err = p.BatchInsert(&ps)
		tools.HasError(err, msg.Failed, 500)
	} else {
		err = p.DeleteActionByPid()
		tools.HasError(err, msg.Failed, 500)
	}*/
	err = p.Update()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}

func DeletePermission(c *gin.Context) {
	var (
		p model.Permission
	)
	id, err := tools.StringToInt(c.Param("permissionId"))
	tools.HasError(err, msg.Failed, -1)
	p.PermissionId = id
	p.UpdatedBy = jwt.GetUserIdStr(c)
	err = p.Delete()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}
