package api

import (
	"basic-antd/init/cache"
	"basic-antd/init/config"
	"basic-antd/init/database"
	"basic-antd/init/global"
	"basic-antd/init/logger"
	"basic-antd/internal/app/router"
	"basic-antd/pkg/casbin"
	"basic-antd/tools"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	configYml string
	port      string
	mode      string
	StartCmd  = &cobra.Command{
		Use:          "start",
		Short:        "Start API server",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "configs/setting.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

func setup() {
	config.Setup(configYml)
	err := logger.Setup()
	if err != nil {
		fmt.Println(err)
	}
	database.Setup()
	global.CasbinEnforcer = casbin.Setup(global.Eloquent)
	cache.Setup()
	// global.CasbinEnforcer = casbin.Setup(global.Eloquent)
	usageStr := `starting api server`
	zap.L().Info(usageStr)
	tools.ClearLimit()
}

func run() error {
	if viper.GetString("setting.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	for _, f := range AppRouters {
		f()
	}
	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: global.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ",
				zap.Error(err),
			)
		}
	}()

	/*go func() {
		// 服务连接
		if config.SSLConfig.Enable {
			if err := srv.ListenAndServeTLS(config.SSLConfig.Pem, config.SSLConfig.KeyStr); err != nil {
				global.Logger.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				global.Logger.Fatal("listen: ", err)
			}
		}
	}()*/
	/*content, _ := ioutil.ReadFile("./assets/start.txt")
	fmt.Println(tools.Red(string(content)))*/
	tip()
	fmt.Println(tools.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", tools.GetLocalHost(), config.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", time.Now().Format("2006-01-02 15:04:05"))
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format("2006-01-02 15:04:05"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:",
			zap.Error(err),
		)
	}
	zap.L().Info("Server exiting")
	return nil
}

func tip() {
	usageStr := `欢迎进入` + tools.Green(`app`+global.Version) + "的世界" + `可以使用` + tools.Red("-h") + `查看命令`
	fmt.Printf("%s \n\n", usageStr)
}
