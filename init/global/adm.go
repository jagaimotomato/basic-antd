package global

import (
	"basic-antd/init/config"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/glog"
	"gorm.io/gorm"
)

var Eloquent *gorm.DB

const Version = "1.0.0"

var (
	Logger        *glog.Logger
	RequestLogger *glog.Logger
)

var (
	Source string
	Driver string
	DBName string
)

var Cfg config.Conf = config.DefaultConfig()
var GinEngine *gin.Engine
var CasbinEnforcer *casbin.SyncedEnforcer
