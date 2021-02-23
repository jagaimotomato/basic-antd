package database

import (
	"basic-antd/init/config"
	global "basic-antd/init/global"
	"basic-antd/tools"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Mysql struct {
}

func Setup() {
	var m = new(Mysql)
	global.Source = m.GetConnect()
	global.Logger.Info(tools.Green(global.Source))
	db, err := sql.Open("mysql", global.Source)
	if err != nil {
		global.Logger.Fatal(tools.Red(m.GetDriver()+"connect error:"), err)
	}
	global.Cfg.SetDb(&config.DBConfig{
		Driver: "mysql",
		DB:     db,
	})
	global.Eloquent, err = m.Open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		global.Logger.Fatal(tools.Red(m.GetDriver()+" connect error :"), err)
	} else {
		global.Logger.Info(tools.Green(m.GetDriver() + " connect success !"))
	}

	if global.Eloquent.Error != nil {
		global.Logger.Fatal(tools.Red("database error :"), global.Eloquent.Error)
	}

	if config.LoggerConfig.EnabledDB {
		global.Eloquent.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	}
}

// 打开数据库连接
func (m *Mysql) Open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}

// 获取数据库连接
func (m *Mysql) GetConnect() string {
	return config.DatabaseConfig.Source
}

func (m *Mysql) GetDriver() string {
	return config.DatabaseConfig.Driver
}
