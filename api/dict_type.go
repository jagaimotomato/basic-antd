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

func GetDictTypeList(c *gin.Context) {
	var d model.DictType
	d.Name = c.Request.FormValue("name")
	d.Type = c.Request.FormValue("type")
	d.DictId, _ = tools.StringToInt(c.Param("dictId"))
	result, count, err := d.GetList()
	tools.HasError(err, msg.Failed, 500)
	app.ListOk(c, result, count, msg.Success)
}

func GetDictTypePage(c *gin.Context) {
	var (
		d         model.DictType
		pageIndex = 1
		pageSize  = 10
		err       error
	)
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = tools.StringToInt(index)
	}
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = tools.StringToInt(size)
	}
	d.DictId, err = tools.StringToInt(c.Request.FormValue("dictId"))
	d.Name = c.Request.FormValue("name")
	d.Type = c.Request.FormValue("type")
	d.Status = c.Request.FormValue("status")
	result, count, err := d.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	app.PaginateOk(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetDictType(c *gin.Context) {
	var d model.DictType
	d.Name = c.Request.FormValue("name")
	d.DictId, _ = tools.StringToInt(c.Param("dictId"))
	result, err := d.Get()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, result, msg.Success)
}

func CreateDictType(c *gin.Context) {
	var d model.DictType
	err := c.MustBindWith(&d, binding.JSON)
	d.CreatedBy = jwt.GetUserIdStr(c)
	tools.HasError(err, msg.Failed, -1)
	err = d.Insert()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, "", msg.Success)
}

func UpdateDictType(c *gin.Context) {
	var d model.DictType
	err := c.MustBindWith(&d, binding.JSON)
	d.UpdatedBy = jwt.GetUserIdStr(c)
	tools.HasError(err, msg.Failed, -1)
	result, err := d.Update()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, result, msg.Success)
}

func DeleteDictType(c *gin.Context) {
	var d model.DictType
	d.UpdatedBy = jwt.GetUserIdStr(c)
	ids := tools.IdStrToIdsGroup("dictId", c)
	err := d.Delete(ids)
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, "", msg.Success)
}
