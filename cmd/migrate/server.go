package migrate

import (
	initialize "basic-antd/init"
	orm "basic-antd/init/global"
	"basic-antd/model"
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var (
	configYml string
	exec      bool
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "initialize the db",
		Example: "blog migrate -c config/settings.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/setting.yml", "Start server with provided configuration file")
}

func setup() {
	initialize.MigrateApplication(configYml)
}

func run() {
	usage := `start init`
	fmt.Println(usage)
	Migrate()
}

func Migrate() {
	fmt.Println("数据库迁移开始")
	err := orm.Eloquent.Debug().AutoMigrate(
		model.Api{},
	)

	if err != nil {
		zap.L().Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	zap.L().Info("register table success")
	fmt.Println("数据库迁移成功")
}
