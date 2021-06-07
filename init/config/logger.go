package config

import "github.com/spf13/viper"

type Logger struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		Level:      cfg.GetString("user"),
		Filename:   cfg.GetString("filename"),
		MaxSize:    cfg.GetInt("maxsize"),
		MaxAge:     cfg.GetInt("maxage"),
		MaxBackups: cfg.GetInt("maxbackups"),
	}
}

var LoggerConfig = new(Logger)
