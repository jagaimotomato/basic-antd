package middleware

import (
	"basic-antd/init/global"
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/response"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthCheck 权限检查中间件
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res bool
		)
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwt.MapClaims)
		roleArr := strings.Split(v["roleName"].(string), ",")
		e := global.CasbinEnforcer
		// msgID := tools.GenerateMsgIdFromContext(c)
		if tools.StrContains("admin", roleArr) {
			res = true
		} else {
			for _, v := range roleArr {
				res, _ = e.Enforce(v, c.Request.URL.Path, c.Request.Method)
				fmt.Println(res, v, c.Request.URL.Path, c.Request.Method)
				if res {
					break
				}
			}
		}

		if res {
			c.Next()
		} else {
			response.Error(c, 403, nil, "您没有该接口的访问权限,请联系管理员")
			c.Abort()
			return
		}
	}
}
