package config

import "github.com/spf13/viper"

type Redis struct {
	Addr     string
	Password string
	DB       int
}

func InitRedis(cfg *viper.Viper) *Redis {
	return &Redis{
		Addr:     cfg.GetString("addr"),
		Password: cfg.GetString("password"),
		DB:       cfg.GetInt("db"),
	}
}

var RedisConfig = new(Redis)
