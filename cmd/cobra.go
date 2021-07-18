package cmd

import (
	"basic-antd/cmd/api"
	"basic-antd/cmd/migrate"
	"basic-antd/init/global"
	"basic-antd/tools"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "response",
	Short:        "response",
	SilenceUsage: true,
	Long:         `application`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(tools.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + tools.Green(`App `+global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	/*rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(config.StartCmd)*/
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
