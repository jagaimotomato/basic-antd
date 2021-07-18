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

func GetRolePage(c *gin.Context) {
	var (
		r         model.Role
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
	r.Name = c.Request.FormValue("name")
	r.RoleKey = c.Request.FormValue("roleKey")
	r.Describe = c.Request.FormValue("describe")
	r.Status = c.Request.FormValue("status")
	result, count, err := r.GetRolePermissionPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	response.Paginate(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetRoleList(c *gin.Context) {
	var r model.Role
	roleList, err := r.GetList()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, roleList, msg.Success)
}

func GetRole(c *gin.Context) {
	var r model.Role
	r.RoleId, _ = tools.StringToInt(c.Param("roleId"))
	role, err := r.Get()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, role, msg.Success)
}

func CreateRole(c *gin.Context) {
	var (
		r  model.Role
		rp model.RolePermission
	)
	err := c.MustBindWith(&r, binding.JSON)
	tools.HasError(err, msg.Failed, -1)
	r.CreatedBy = jwt.GetUserIdStr(c)
	roleId, err := r.Insert()
	tools.HasError(err, msg.Failed, 500)
	if len(r.PermissionIds) > 0 {
		err = rp.Insert(roleId, r.PermissionIds)
		tools.HasError(err, msg.Failed, 500)
	}
	if len(r.Apis) > 0 {
		err = rp.InsertCasbin(roleId, r.Apis)
		tools.HasError(err, msg.Failed, 500)
	}
	response.Success(c, "", msg.Success)
}

func UpdateRole(c *gin.Context) {
	var (
		r       model.Role
		rp      model.RolePermission
		roleIds = make([]int, 1)
	)
	err := c.MustBindWith(&r, binding.JSON)
	tools.HasError(err, msg.Failed, -1)
	r.UpdatedBy = jwt.GetUserIdStr(c)
	err = r.Update()
	tools.HasError(err, msg.Failed, 500)
	roleIds[0] = r.RoleId
	err = rp.DeleteRolePermission(roleIds)
	tools.HasError(err, msg.Failed, -1)
	if len(r.PermissionIds) > 0 {
		err = rp.Insert(r.RoleId, r.PermissionIds)
		tools.HasError(err, msg.Failed, 500)
	}
	if len(r.Apis) > 0 {
		err = rp.InsertCasbin(r.RoleId, r.Apis)
		tools.HasError(err, msg.Failed, 500)
	}
	response.Success(c, "", msg.Success)
}

func UpdateRoleStatus(c *gin.Context) {
	var r model.Role
	r.Status = c.Request.FormValue("status")
	r.RoleId, _ = tools.StringToInt(c.Request.FormValue("roleId"))
	r.UpdatedBy = jwt.GetUserIdStr(c)
	err := r.Update()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}

func DeleteRole(c *gin.Context) {
	var (
		r  model.Role
		rp model.RolePermission
	)
	r.UpdatedBy = jwt.GetUserIdStr(c)
	roleIds := tools.IdStrToIdsGroup("roleId", c)
	err := r.Delete(roleIds)
	if err != nil {
		response.Error(c, 400, err, "")
		return
	} else {
		err = rp.DeleteRolePermission(roleIds)
		tools.HasError(err, msg.Failed, 500)
		response.Success(c, "", msg.Success)
	}
}
