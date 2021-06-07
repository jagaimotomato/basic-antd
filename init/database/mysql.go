package database

import (
	"basic-antd/init/config"
	"basic-antd/init/global"
	"basic-antd/tools"
	"database/sql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Setup() {
	source := config.DatabaseConfig.Source
	// fmt.Println(source, "******")
	zap.L().Info(tools.Green("数据库source"),
		zap.String("source", source),
	)
	db, err := sql.Open("mysql", source)
	if err != nil {
		global.Logger.Fatal(tools.Red("打开mysql失败"),
			zap.Error(err),
		)
	}
	global.Eloquent, err = open(db, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Fatal(tools.Red("gorm connect error"),
			zap.Error(err),
		)
	} else {
		zap.L().Info(tools.Green("gorm connect success"))
	}
}

func open(db *sql.DB, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{Conn: db}), cfg)
}
