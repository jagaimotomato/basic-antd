package logger

import (
	"basic-antd/init/config"
	global "basic-antd/init/global"
	"basic-antd/tools"
	"github.com/gogf/gf/os/glog"
)

var (
	Logger        *glog.Logger
	RequestLogger *glog.Logger
)

func Setup() {
	Logger = glog.New()
	_ = Logger.SetPath(config.LoggerConfig.Path + "/bus")
	Logger.SetStdoutPrint(config.LoggerConfig.EnabledBUS && config.LoggerConfig.Stdout)
	Logger.SetFile("bus-{Ymd}.log")
	_ = Logger.SetLevelStr(config.LoggerConfig.Level)

	RequestLogger = glog.New()
	_ = RequestLogger.SetPath(config.LoggerConfig.Path + "/request")
	RequestLogger.SetStdoutPrint(false)
	RequestLogger.SetFile("access-{Ymd}.log")
	_ = RequestLogger.SetLevelStr(config.LoggerConfig.Level)

	Logger.Info(tools.Green("Logger init success!"))

	global.Logger = Logger.Line()
	global.RequestLogger = RequestLogger.Line()
}
