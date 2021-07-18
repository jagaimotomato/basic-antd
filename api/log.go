package api

import (
	"basic-antd/model"
	"basic-antd/pkg/response"
	"basic-antd/pkg/response/msg"
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
	if index := c.Request.FormValue("page_index"); index != "" {
		pageIndex, _ = tools.StringToInt(index)
	}
	if size := c.Request.FormValue("page_size"); size != "" {
		pageSize, _ = tools.StringToInt(size)
	}
	result, count, err := l.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	response.Paginate(c, result, count, pageIndex, pageSize, msg.Success)
}

func GetOperationLogPage(c *gin.Context) {
	var (
		o         model.OperationLog
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
	result, count, err := o.GetPage(pageIndex, pageSize)
	tools.HasError(err, msg.Failed, 500)
	response.Paginate(c, result, count, pageIndex, pageSize, msg.Success)
}
