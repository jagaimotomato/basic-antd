package router

import (
	"basic-antd/init/global"
	"basic-antd/internal/app/middleware"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() {
	r := global.Engine
	middleware.InitMiddleware(r)
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "Jwt init error", 500)
	initAppRouter(r, authMiddleware)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong"+fmt.Sprint(time.Now().Unix()))
	})
}
