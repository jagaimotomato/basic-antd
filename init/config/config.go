package config

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	cfgDatabase    *viper.Viper
	cfgApplication *viper.Viper
	cfgJwt         *viper.Viper
	cfgRedis       *viper.Viper
	cfgMail        *viper.Viper
	cfgLogger      *viper.Viper
)

func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgDatabase = viper.Sub("setting.database")
	if cfgDatabase == nil {
		panic("No found setting.database in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("setting.application")
	if cfgApplication == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("setting.jwt")
	if cfgJwt == nil {
		panic("No found settings.jwt in the configuration")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgLogger = viper.Sub("setting.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLogger(cfgLogger)

	cfgRedis = viper.Sub("setting.redis")
	if cfgRedis == nil {
		panic("No found settings.redis in the configuration")
	}
	RedisConfig = InitRedis(cfgRedis)

	cfgMail = viper.Sub("setting.mail")
	if cfgMail == nil {
		panic("No found settings.mail in the configuration")
	}
	MailConfig = InitMail(cfgMail)
}

type Config struct {
	saas   bool
	dbs    map[string]*DBConfig
	db     *DBConfig
	engine http.Handler
}

type DBConfig struct {
	Driver string
	DB     *sql.DB
}

// SetDbs 设置对应key的db
func (c *Config) SetDbs(key string, db *DBConfig) {
	c.dbs[key] = db
}

// GetDbs 获取所有map里的db数据
func (c *Config) GetDbs() map[string]*DBConfig {
	return c.dbs
}

// GetDbByKey 根据key获取db
func (c *Config) GetDbByKey(key string) *DBConfig {
	return c.dbs[key]
}

// SetDb 设置单个db
func (c *Config) SetDb(db *DBConfig) {
	c.db = db
}

// GetDb 获取单个db
func (c *Config) GetDb() *DBConfig {
	return c.db
}

// SetEngine 设置路由引擎
func (c *Config) SetEngine(engine http.Handler) {
	c.engine = engine
}

// GetEngine 获取路由引擎
func (c *Config) GetEngine() http.Handler {
	return c.engine
}

func DefaultConfig() *Config {
	return &Config{}
}
