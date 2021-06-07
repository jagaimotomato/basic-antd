package api

import (
	"basic-antd/internal/app/model"
	"basic-antd/pkg/app"
	"basic-antd/pkg/app/msg"
	"basic-antd/pkg/jwt"
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetUserPage(c *gin.Context) {
	var (
		user      model.User
		pageSize  = 10
		pageIndex = 1
		err       error
	)
	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize, err = tools.StringToInt(size)
	}
	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex, err = tools.StringToInt(index)
	}
	// fmt.Println(pageIndex, pageSize, "************************")
	user.Username = c.Request.FormValue("username")
	user.Phone = c.Request.FormValue("phone")
	user.Name = c.Request.FormValue("name")
	user.DepartmentId, err = tools.StringToInt(c.Request.FormValue("departmentId"))
	result, count, err := user.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", 500)
	app.PaginateOk(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetUser(c *gin.Context) {
	var user model.User
	userId, _ := tools.StringToInt(c.Param("userId"))
	user, err := user.GetUserById(userId)
	tools.HasError(err, "", 500)
	app.Success(c, user, msg.Success)
}

func CreateUser(c *gin.Context) {
	var (
		user model.User
		ur   model.UserRole
		du   model.DepartmentUser
	)
	err := c.MustBindWith(&user, binding.JSON)
	tools.HasError(err, "数据格式不正确", 500)
	user.CreatedBy = jwt.GetRealName(c)
	// 生成加盐值
	user.Password, user.Salt, err = tools.GeneratePassword(user.Password)
	tools.HasError(err, "加盐失败", 500)
	userId, err := user.Insert()
	tools.HasError(err, "", 500)
	if len(user.RoleIds) > 0 {
		err = ur.Insert(userId, user.RoleIds)
		tools.HasError(err, msg.Failed, 500)
	}
	if len(user.DepartmentIds) > 0 {
		err = du.Insert(userId, user.DepartmentIds)
		tools.HasError(err, msg.Failed, 500)
	}
	app.Success(c, "", msg.Success)
}

func UpdateUser(c *gin.Context) {
	var (
		user    model.User
		ur      model.UserRole
		du      model.DepartmentUser
		userIds = make([]int, 1)
	)
	err := c.MustBindWith(&user, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	user.UpdatedBy = jwt.GetRealName(c)
	result, err := user.Update()
	tools.HasError(err, "", 500)
	userIds[0] = user.UserId
	err = ur.DeleteUserRole(userIds)
	tools.HasError(err, msg.Failed, -1)
	if len(user.RoleIds) > 0 {
		err = ur.Insert(user.UserId, user.RoleIds)
		tools.HasError(err, msg.Failed, -1)
		err = du.Insert(user.UserId, user.DepartmentIds)
		tools.HasError(err, msg.Failed, -1)
	}
	app.Success(c, result, msg.Success)
}

func UpdateUserStatus(c *gin.Context) {
	var u model.User
	u.Status = c.Request.FormValue("status")
	u.UserId, _ = tools.StringToInt(c.Request.FormValue("userId"))
	u.UpdatedBy = jwt.GetUserIdStr(c)
	result, err := u.Update()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, result, msg.Success)
}

func DeleteUser(c *gin.Context) {
	var (
		user model.User
		ur   model.UserRole
		du   model.DepartmentUser
	)
	user.UpdatedBy = jwt.GetRealName(c)
	//fmt.Println(jwt.GetRealName(c), jwt.GetRoleId(c), jwt.GetRoleName(c), jwt.GetUserName(c), jwt.GetUserId(c))
	Ids := tools.IdStrToIdsGroup("userId", c)
	err := user.BatchDelete(Ids)
	if err != nil {
		app.ResponseError(c, 400, err, "")
		return
	} else {
		err = ur.DeleteUserRole(Ids)
		tools.HasError(err, msg.Failed, 500)
		err = du.DeleteDepartmentUser(Ids)
		tools.HasError(err, msg.Failed, 500)
		app.Success(c, "", msg.Success)
	}
}
