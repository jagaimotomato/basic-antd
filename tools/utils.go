package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"log"
	"runtime"
	"strconv"
)

func HasError(err error, msg string, code int) {
	if err != nil {
		statusCode := 200
		if msg == "" {
			msg = err.Error()
		}
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%v error: %#v", file, line, err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// 生成http请求的requestId
func GenerateMsgIdFromContext(c *gin.Context) string {
	var msgId string
	data, ok := c.Get("msgID")
	if !ok {
		msgId = uuid.New().String()
		c.Set("msgID", msgId)
		return msgId
	}
	msgId = cast.ToString(data)
	return msgId
}
