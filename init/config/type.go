package config

import (
	"basic-antd/pkg/logger"
	"net/http"
)

type Conf interface {
	//多db设置，⚠ SetDbs 不允许并发，根据业务，例如 app分库，host分库
	SetDbs(key string, db *DBConfig)
	GetDbs() map[string]*DBConfig
	GetDbByKey(key string) *DBConfig
	GetSaas() bool
	SetSaas(bool)

	//单库业务实现这两个接口
	SetDb(db *DBConfig)
	GetDb() *DBConfig

	//使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	//使用定义的logger，参考来源go-micro
	SetLogger(logger logger.Logger)
	GetLogger() logger.Logger
}
