package database

import (
	"basic-antd/init/config"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	switch getDriver() {
	case "mysql":
		return connectMysql()
	default:
		return connectMysql()
	}
}

func getConnect() string {
	return config.DatabaseConfig.Source
}

func getDriver() string {
	return config.DatabaseConfig.Driver
}
