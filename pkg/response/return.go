package response

import (
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	res.RequestId = tools.GenerateMsgIdFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.Error(code))
}

func Success(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Result = data
	if msg != "" {
		res.Msg = msg
	}
	res.RequestId = tools.GenerateMsgIdFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res.Success())
}

func Paginate(c *gin.Context, result interface{}, count int64, pageIndex, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	Success(c, res, msg)
}

func ListSuccess(c *gin.Context, result interface{}, count int64, msg string) {
	var res List
	res.List = result
	res.Count = count
	Success(c, res, msg)
}

func Custom(c *gin.Context, data gin.H) {
	c.AbortWithStatusJSON(http.StatusOK, data)
}
