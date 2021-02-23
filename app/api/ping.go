package api

import (
	"basic-antd/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func AuthPong(c *gin.Context) {
	roleIds := jwt.GetRoleIdStr(c)
	roleName := jwt.GetRoleName(c)
	c.JSON(http.StatusOK, gin.H{
		"success":  "ok",
		"roleIds":  roleIds,
		"roleName": roleName,
	})
}
