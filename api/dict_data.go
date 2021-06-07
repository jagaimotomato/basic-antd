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

func GetDictDataList(c *gin.Context) {
	var d model.DictData
	d.Type = c.Param("type")
	result, count, err := d.GetList()
	tools.HasError(err, msg.Failed, 500)
	app.ListOk(c, result, count, msg.Success)
}

func GetDictDataPage(c *gin.Context) {
	var (
		d         model.DictData
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

	d.Label = c.Request.FormValue("label")
	d.Type = c.Request.FormValue("type")
	d.Value = c.Request.FormValue("value")
	d.Status = c.Request.FormValue("status")
	d.DictCode, err = tools.StringToInt(c.Request.FormValue("dictCode"))
	result, count, err := d.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	app.PaginateOk(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetDictData(c *gin.Context) {
	var d model.DictData
	d.DictCode, _ = tools.StringToInt(c.Param("dictCode"))
	result, err := d.Get()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, result, msg.Success)
}

func CreateDictData(c *gin.Context) {
	var d model.DictData
	err := c.MustBindWith(&d, binding.JSON)
	d.CreatedBy = jwt.GetUserIdStr(c)
	tools.HasError(err, msg.Failed, -1)
	err = d.Insert()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, "", msg.Success)
}

func UpdateDictData(c *gin.Context) {
	var d model.DictData
	err := c.MustBindWith(&d, binding.JSON)
	d.UpdatedBy = jwt.GetUserIdStr(c)
	tools.HasError(err, msg.Failed, -1)
	result, err := d.Update()
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, result, msg.Success)
}

func DeleteDictData(c *gin.Context) {
	var d model.DictData
	d.UpdatedBy = jwt.GetUserIdStr(c)
	ids := tools.IdStrToIdsGroup("dictCode", c)
	err := d.Delete(ids)
	tools.HasError(err, msg.Failed, 500)
	app.Success(c, "", msg.Success)
}
