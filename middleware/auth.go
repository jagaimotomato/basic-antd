package middleware

import (
	"basic-antd/init/config"
	handler2 "basic-antd/middleware/handler"
	"basic-antd/pkg/jwt"
	"time"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.ApplicationConfig.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
		/*if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}*/
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",                     // JWT标识
		Key:             []byte(config.JwtConfig.Secret), // 服务端密钥
		Timeout:         timeout,                         //  token 过期时间
		MaxRefresh:      time.Hour,                       // token 更新时间
		Authenticator:   handler2.Authenticator,          // 在登录接口中使用的验证方法，并返回验证成功后的用户对象。
		Authorizator:    handler2.Authorizator,           // 登录后其他接口验证传入的 token 方法
		PayloadFunc:     handler2.PayloadFunc,            // 添加额外业务相关的信息
		Unauthorized:    handler2.Unauthoried,            // 验证失败后设置错误信息
		IdentityHandler: handler2.IdentityHandler,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt", // 设置 token 获取位置，一般默认在头部的 Authorization 中，或者 query的 token 字段，cookie 中的 jwt 字段
		TokenHeadName:   "Bearer",                                           // Header中 token 的头部字段，默认常用名称 Bearer
		TimeFunc:        time.Now,
	})
}
