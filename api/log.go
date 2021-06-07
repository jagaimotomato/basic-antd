package api

import (
	"basic-antd/internal/app/model"
	"basic-antd/pkg/app"
	"basic-antd/pkg/app/msg"
	"basic-antd/tools"
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
