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

func GetDepartment(c *gin.Context) {
	var d model.Department
	d.DepartmentId, _ = tools.StringToInt(c.Param("department_id"))
	d, err := d.Get()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, d, msg.Success)
}

func GetDepartmentPage(c *gin.Context) {
	var (
		d         model.Department
		pageIndex = 1
		pageSize  = 10
		err       error
	)
	if index := c.Request.FormValue("page_index"); index != "" {
		pageIndex, _ = tools.StringToInt(index)
	}
	if size := c.Request.FormValue("page_size"); size != "" {
		pageSize, _ = tools.StringToInt(size)
	}
	result, count, err := d.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	response.Paginate(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetDepartmentTree(c *gin.Context) {
	var (
		d          model.Department
		department []model.Department
	)
	d.Name = c.Request.FormValue("name")
	department, err := d.GetDepartmentTree()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, department, msg.Success)
}

func CreateDepartment(c *gin.Context) {
	var d model.Department
	err := c.MustBindWith(&d, binding.JSON)
	tools.HasError(err, msg.BindingFailed, -1)
	d.CreatedBy = jwt.GetUserIdStr(c)
	err = d.Insert()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}

func UpdateDepartment(c *gin.Context) {
	var d model.Department
	err := c.MustBindWith(&d, binding.JSON)
	tools.HasError(err, msg.BindingFailed, -1)
	err = d.Update()
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}

func DeleteDepartment(c *gin.Context) {
	var d model.Department
	ids := tools.IdStrToIdsGroup("department_id", c)
	err := d.Delete(ids)
	tools.HasError(err, msg.Failed, 500)
	response.Success(c, "", msg.Success)
}
