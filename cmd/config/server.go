package config

import (
	"basic-antd/init/config"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configYml string
	mode      string
	StartCmd  = &cobra.Command{
		Use:   "config",
		Short: "Get Application Config Info",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided config")
}

func run() {
	config.Setup(configYml)
	application, err := json.MarshalIndent(config.ApplicationConfig, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("application", string(application))
	jwt, err := json.MarshalIndent(config.JwtConfig, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("application", string(jwt))

	database, err := json.MarshalIndent(config.DatabaseConfig, "", "   ") //转换成JSON返回的是byte[]
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("database:", string(database))

	redis, err := json.MarshalIndent(config.RedisConfig, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("redis:", string(redis))

	loggerConfig, err := json.MarshalIndent(config.LoggerConfig, "", "   ") //转换成JSON返回的是byte[]
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("logger:", string(loggerConfig))
}
