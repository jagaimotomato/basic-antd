package config

import "github.com/spf13/viper"

type Mail struct {
	User     string
	Password string
	Host     string
	Port     int
	Alias    string
}

func InitMail(cfg *viper.Viper) *Mail {
	return &Mail{
		User:     cfg.GetString("user"),
		Password: cfg.GetString("password"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetInt("port"),
		Alias:    cfg.GetString("alias"),
	}
}

var MailConfig = new(Mail)
