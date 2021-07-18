package init

import (
	"basic-antd/init/cache"
	"basic-antd/init/config"
	"basic-antd/init/database"
	"basic-antd/init/global"
	"basic-antd/init/logger"
	"basic-antd/pkg/casbin"
)

func ApiApplication(path string) {
	config.Setup(path)
	err := logger.Setup()
	if err != nil {
		panic(err)
	}
	global.Eloquent = database.Setup()
	global.Rdb = cache.RedisSetup()
	global.CasbinEnforcer = casbin.Setup(global.Eloquent)
}

func MigrateApplication(path string) {
	config.Setup(path)
	err := logger.Setup()
	if err != nil {
		panic(err)
	}
	global.Eloquent = database.Setup()
}
