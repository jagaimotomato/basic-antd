package app

import (
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, code int, err error, msg string) {
	var res Response
	if err != nil {
		res.Message = err.Error()
	}
	if msg != "" {
		res.Message = msg
	}
	res.RequestId = tools.GenerateMsgIdFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnError(code))
}

func Success(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Result = data
	if msg != "" {
		res.Message = msg
	}
	res.RequestId = tools.GenerateMsgIdFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.ReturnSuccess())
}

func PaginateOk(c *gin.Context, result interface{}, count int64, pageIndex, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	Success(c, res, msg)
}

func ListOk(c *gin.Context, result interface{}, count int64, msg string) {
	var res List
	res.List = result
	res.Count = count
	Success(c, res, msg)
}

func Custom(c *gin.Context, data gin.H) {
	c.AbortWithStatusJSON(http.StatusOK, data)
}
