package config

import (
	"errors"
	"fmt"

	"github.com/kukkar/common-golang/pkg/components/mongodb"
	"github.com/kukkar/common-golang/pkg/config"
	// "github.com/kukkar/common-golang/pkg/factory/cache"
	// "github.com/kukkar/common-golang/pkg/factory/sql"
)

type AppConfig struct {
	Mongo        *mongodb.MDBConfig `json:"Mongo,omitempty"`
	ImageStorage struct {
		Use   string
		Local struct {
			Path string
		}
		S3 struct {
			BucketName string
			Region     string
			LogLevel   int
		}
	}
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

	m["Mongo.URL"] = "TIGERHALL_MONGO_URL"
	return m
}
