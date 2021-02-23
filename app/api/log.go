package api

import (
	"basic-antd/app/model"
	"basic-antd/tools"
	"basic-antd/tools/app"
	"basic-antd/tools/app/msg"
	"github.com/gin-gonic/gin"
)

func GetLoginLogPage(c *gin.Context) {
	var (
		l         model.LoginLog
		pageIndex = 1
		pageSize  = 10
		err       error
	)
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, _ = tools.StringToInt(index)
	}
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = tools.StringToInt(size)
	}
	result, count, err := l.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	app.PaginateOk(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetOperationLogPage(c *gin.Context) {
	var (
		o         model.OperationLog
		pageIndex = 1
		pageSize  = 10
		err       error
	)
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, _ = tools.StringToInt(index)
	}
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = tools.StringToInt(size)
	}
	result, count, err := o.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	app.PaginateOk(c, result, count, pageIndex, pageSize, msg.Success)
}
