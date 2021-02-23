package router

import (
	"basic-antd/app/middleware"
	v1 "basic-antd/app/router/v1"
	"basic-antd/init/global"
	"basic-antd/tools"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func InitRouter() {
	var r *gin.Engine
	h := global.Cfg.GetEngine()
	if h == nil {
		h = gin.New()
		global.Cfg.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	middleware.InitMiddleware(r)
	authMiddle, err := middleware.AuthInit()
	tools.HasError(err, "Jwt init error", 500)
	v1.InitBaseRouter(r, authMiddle)
	//return r
}
