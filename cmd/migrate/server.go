package migrate

import (
	"basic-antd/app/model"
	"basic-antd/init/config"
	"basic-antd/init/database"
	"basic-antd/init/global"
	"basic-antd/init/logger"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configYml string
	exec      bool
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "initialize the db",
		Example: "blog migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func run() {
	usage := `start init`
	fmt.Println(usage)
	config.Setup(configYml)
	logger.Setup()
	database.Setup()
	fmt.Println("数据库迁移开始")
	err := migrate()
	if err != nil {
		global.Logger.Fatalf("数据库迁移失败，错误：", err)
	}
	usage = `数据库初始化成功`
	fmt.Println(usage)
}

func migrate() error {
	if config.DatabaseConfig.Driver == "mysql" {
		global.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	}

	err := global.Eloquent.Debug().AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.RolePermission{},
		model.UserRole{},
		&model.Permission{},
		&model.Department{},
		&model.DepartmentUser{},
		&model.DictType{},
		&model.DictData{},
		&model.OperationLog{},
		&model.LoginLog{},
	)
	return err
}
