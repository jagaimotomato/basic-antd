package router

import (
	"basic-antd/init/global"
	middleware2 "basic-antd/middleware"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() {
	r := global.Engine
	middleware2.InitMiddleware(r)
	authMiddleware, err := middleware2.AuthInit()
	tools.HasError(err, "Jwt init error", 500)
	initAppRouter(r, authMiddleware)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong"+fmt.Sprint(time.Now().Unix()))
	})
}
