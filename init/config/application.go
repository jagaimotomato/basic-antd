package config

import "github.com/spf13/viper"

type Application struct {
	ReadTimeout  int
	WriteTimeout int
	Host         string
	Port         string
	Name         string
	Mode         string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:  cfg.GetInt("readTimeout"),
		WriteTimeout: cfg.GetInt("writeTimeout"),
		Host:         cfg.GetString("host"),
		Port:         cfg.GetString("port"),
		Name:         cfg.GetString("name"),
		Mode:         cfg.GetString("mode"),
	}
}

var ApplicationConfig = new(Application)
