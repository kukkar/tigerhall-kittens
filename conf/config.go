package config

import (
	"errors"
	"fmt"

	"github.com/kukkar/common-golang/pkg/config"
	// "github.com/kukkar/common-golang/pkg/factory/cache"
	// "github.com/kukkar/common-golang/pkg/factory/sql"
)

type AppConfig struct {
	//	MySql              *sql.MysqlConfig   `json:"Mysql"`
	//	Cache              *cache.CacheConfig `json:"Cache"`
	UseClientValidator bool   `json:"UseClientValidator"`
	SentryDSN          string `json:"SentryDSN"`
}

func GetAppConfig() (*AppConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig
	appConfig, ok := c.(*AppConfig)
	if !ok {
		msg := fmt.Sprintf("Example APP Config Not correct %+v", c)
		return nil, errors.New(msg)
	}
	return appConfig, nil
}

func GetGlobalConfig() (*config.AppConfig, error) {
	return config.GlobalAppConfig, nil
}

func EnvUpdateMap() map[string]string {
	m := make(map[string]string)

	m["Mysql.User"] = "tigerhall-kittens_MYSQL_USER"
	m["Mysql.Password"] = "tigerhall-kittens_MYSQL_PASSWORD"
	m["Mysql.DbName"] = "tigerhall-kittens_MYSQL_DBNAME"
	m["Mysql.MaxOpenConnections"] = "tigerhall-kittens_MYSQL_MAXOPENCONNECTIONS"
	m["Mysql.MaxIdleConnections"] = "tigerhall-kittens_MYSQL_MAXIDLECONNECTIONS"
	m["Mysql.DefaultTimeZone"] = "tigerhall-kittens_MYSQL_DEFAULTTIMEZONE"
	m["Mysql.Host"] = "tigerhall-kittens_MYSQL_HOST"
	m["Mysql.Port"] = "tigerhall-kittens_MYSQL_PORT"
	m["Cache.Use"] = "tigerhall-kittens_CACHE_USE"
	m["Cache.Redis.Addr"] = "tigerhall-kittens_CACHE_REDIS_ADDRESS"
	m["Cache.Redis.PoolSize"] = "tigerhall-kittens_CACHE_REDIS_POOLSIZE"

	m["SentryDSN"] = "tigerhall-kittens_SENTRY_DSN"
	return m
}
