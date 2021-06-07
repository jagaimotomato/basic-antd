package config

import "github.com/spf13/viper"

type Database struct {
	Driver string
	Source string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Driver: cfg.GetString("driver"),
		Source: cfg.GetString("source"),
	}
}

var DatabaseConfig = new(Database)