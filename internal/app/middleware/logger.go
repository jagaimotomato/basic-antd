package middleware

import (
	"basic-antd/internal/app/model"
	"basic-antd/pkg/jwt"
	"basic-antd/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginzap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				zap.L().Error(e)
			}
		} else {
			zap.L().Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
		}
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(
						c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func OperationLogger() gin.HandlerFunc {
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

		if c.Request.Method != "GET" && c.Request.Method != "OPTIONS" {
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
