package middleware

import (
	"basic-antd/init/global"
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/logger"
	"basic-antd/tools"
	"basic-antd/tools/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 权限检查中间件
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res bool
			err error
		)
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwt.MapClaims)
		roleArr := strings.Split(v["roleName"].(string), ",")
		e := global.CasbinEnforcer
		msgID := tools.GenerateMsgIdFromContext(c)
		if tools.StrContains("admin", roleArr) {
			res = true
			logger.Infof("msgID[%s] info:%s method:%s path:%s", msgID, roleArr, c.Request.Method, c.Request.URL.Path)
		} else {
			for _, v := range roleArr {
				fmt.Println(v, c.Request.URL.Path, c.Request.Method)
				res, err = e.Enforce(v, c.Request.URL.Path, c.Request.Method)
				if err != nil {
					logger.Errorf("msgID[%s] error:%s method:%s path:%s", msgID, err, c.Request.Method, c.Request.URL.Path)
					app.ResponseError(c, 500, err, "")
				}
				if res {
					break
				}
			}
		}

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    403,
				"message": "您没有该接口的访问权限,请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
