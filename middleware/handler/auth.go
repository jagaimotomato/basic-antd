package handler

import (
	"basic-antd/model"
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/response"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(model.User)
		roles, roleIds := u.GetRoleStr()
		//r, _ := v["role"].(model.Role)
		return jwt.MapClaims{
			jwt.IdentityKey: u.UserId,
			jwt.Username:    u.Username,
			jwt.RoleId:      roleIds,
			jwt.RoleName:    roles, //roleName
			jwt.RealName:    u.RealName,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"Username":    claims["username"],
		"RoleIds":     claims["roleId"],
		"RoleName":    claims["roleName"],
		"RealName":    claims["realName"],
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var (
		loginForm model.Login
		status    = "1"
		msg       = "登录成功"
		username  = ""
	)
	if err := c.ShouldBind(&loginForm); err != nil {
		username = loginForm.Username
		msg = "数据解析失败"
		status = "0"
		LoginLogoutToDb(c, status, msg, username)
		return nil, jwt.ErrMissingLoginValues
	}
	username = loginForm.Username
	user, flag, err := loginForm.GetUser()
	switch flag {
	case 0:
		LoginLogoutToDb(c, status, msg, username)
		return map[string]interface{}{"user": user}, nil
	case 1:
		if err != nil {
			zap.L().Error(err.Error())
		}
		msg = "用户不存在"
		status = "0"
		LoginLogoutToDb(c, status, msg, username)
		return nil, jwt.ErrFailedUsername
	case 2:
		if err != nil {
			zap.L().Error(err.Error())
		}
		msg = "用户被停用"
		status = "0"
		LoginLogoutToDb(c, status, msg, username)
		return nil, jwt.ErrFailedUserStatus
	case 3:
		if err != nil {
			zap.L().Error(err.Error())
		}
		msg = "密码错误"
		status = "0"
		LoginLogoutToDb(c, status, msg, username)
		return nil, jwt.ErrFailedPassword
	}
	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(model.User)
		roles, roleIds := u.GetRoleStr()
		c.Set("userId", u.UserId)
		c.Set("username", u.Username)
		c.Set("roleName", roles)
		c.Set("roleIds", roleIds)
		c.Set("realName", u.RealName)
		return true
	}
	return false
}

func Unauthoried(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

func Logout(c *gin.Context) {
	status := "1"
	msg := "退出登录"
	username := jwt.GetUserName(c)
	LoginLogoutToDb(c, status, msg, username)
	response.Success(c, "", "退出成功")
}

func LoginLogoutToDb(c *gin.Context, status, msg, username string) {
	var log model.LoginLog
	userAgent := user_agent.New(c.Request.UserAgent())
	log.Ip = c.ClientIP()
	log.Username = username
	log.Location = tools.GetLocation(c.ClientIP())
	log.LoginTime = time.Now()
	log.Status = status
	log.Remark = c.Request.UserAgent()
	browserName, version := userAgent.Browser()
	log.Browser = browserName + version
	log.Os = userAgent.OS()
	log.Msg = msg
	log.Platform = userAgent.Platform()
	if userAgent.Mobile() {
		log.Mobile = "1"
	} else {
		log.Mobile = "0"
	}
	err := log.Insert()
	if err != nil {
		fmt.Println(err)
	}
}
