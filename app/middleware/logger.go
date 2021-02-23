package middleware

import (
	"basic-antd/app/model"
	"basic-antd/init/config"
	"basic-antd/init/global"
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/logger"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		duration := end.Sub(start)
		// 请求方式
		requestMethod := c.Request.Method
		// 请求路由
		requestUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		log := map[string]interface{}{
			"statusCode":    statusCode,
			"duration":      duration,
			"clientIP":      clientIP,
			"requestMethod": requestMethod,
			"requestUri":    requestUri,
		}
		logger.Info(log)
		global.RequestLogger.Info(log)
		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" && config.LoggerConfig.EnabledDB {
			SetDBOperLog(c, clientIP, requestUri, requestMethod, statusCode, duration)
		}
	}
}

func SetDBOperLog(c *gin.Context, clientIP, requestUri, requestMethod string, statusCode int, duration time.Duration) {
	var (
		permission   model.Permission
		operationLog model.OperationLog
	)
	permission.Path = requestUri
	permission.Actions = requestMethod
	permissionList, _ := permission.GetList()
	operationLog.OperIp = clientIP
	operationLog.OperLocation = tools.GetLocation(clientIP)
	operationLog.Code = tools.IntToString(statusCode)
	operationLog.OperName = jwt.GetRealName(c)
	operationLog.RequestMethod = requestMethod
	operationLog.OperUrl = requestUri
	if strings.Contains(requestUri, "/login") {
		operationLog.BusinessType = "10"
		operationLog.Title = "用户登录"
		operationLog.OperName = "-"
	} else if strings.Contains(requestUri, "/logout") {
		operationLog.BusinessType = "11"
	} else {
		switch requestMethod {
		case "POST":
			operationLog.BusinessType = "1"
			break
		case "PUT":
			operationLog.BusinessType = "2"
			break
		case "DELETE":
			operationLog.BusinessType = "3"
			break
		}
	}
	operationLog.Method = requestMethod
	if len(permissionList) > 0 {
		operationLog.Title = permissionList[0].Title
	}
	body, _ := c.Get("body")
	operationLog.OperParam, _ = tools.StructToJsonStr(body)
	operationLog.CreatedBy = jwt.GetUserIdStr(c)
	operationLog.OperTime = time.Now()
	operationLog.DurationTime = (duration).String()
	operationLog.UserAgent = c.Request.UserAgent()
	if c.Err() != nil {
		operationLog.Status = "0"
	} else {
		operationLog.Status = "1"
	}
	err := operationLog.Insert()
	if err != nil {
		fmt.Println(err)
	}
}
