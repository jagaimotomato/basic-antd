package jwt

import (
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) int {
	claims := ExtractClaims(c)
	if claims["identity"] != nil {
		return int(claims["identity"].(float64))
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少identity")
	return 0
}

func GetUserIdStr(c *gin.Context) string {
	claims := ExtractClaims(c)
	if claims["identity"] != nil {
		return tools.Int64ToString(int64((claims["identity"]).(float64)))
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少identity")
	return ""
}

func GetUserName(c *gin.Context) string {
	claims := ExtractClaims(c)
	if claims["username"] != nil {
		return (claims["username"]).(string)
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少username")
	return ""
}

func GetRealName(c *gin.Context) string {
	claims := ExtractClaims(c)
	if claims["realName"] != nil {
		return (claims["realName"]).(string)
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少realName")
	return ""
}

func GetRoleName(c *gin.Context) string {
	claims := ExtractClaims(c)
	if claims["roleName"] != nil {
		return (claims["roleName"]).(string)
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少roleName")
	return ""
}

func GetRoleId(c *gin.Context) int {
	claims := ExtractClaims(c)
	if claims["roleId"] != nil {
		return int((claims["roleId"]).(float64))
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少roleId")
	return 0
}

func GetRoleIdStr(c *gin.Context) string {
	claims := ExtractClaims(c)
	if claims["roleId"] != nil {
		return (claims["roleId"]).(string)
	}
	fmt.Println(tools.GetCurrentTimeStr() + "[WARING]" + c.Request.Method + " " + c.Request.URL.Path + "缺少roleId Str")
	return ""
}
