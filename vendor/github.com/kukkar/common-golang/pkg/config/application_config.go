package config

import (
	"encoding/json"
	"fmt"

	"github.com/kukkar/common-golang/pkg/logger"
)

// AppConfig will contain all the app related config data which should be provided at the start of the app
type AppConfig struct {
	AppName           string        `json:"AppName"`
	AppVersion        string        `json:"AppVersion"`
	ServerHost        string        `json:"ServerHost"`
	ServerPort        string        `json:"ServerPort"`
	Environment       string        `json:"Environment"`
	LogConfFile       string        `json:"LogConfFile"`
	LogConfig         logger.Config `json:"LogConfig"`
	SuperKey          string        `json:"SuperKey"`
	ApplicationConfig interface{}   `json:"ApplicationConfig"`
}

func (this *AppConfig) String() string {
	str, err := json.Marshal(this)
	if err != nil {
		return "Could NOT Marshal APP config."
	}
	return string(str)
}

func (this *AppConfig) ShowConfig() string {

	s := fmt.Sprintf(
		"AppName      := %s\n"+
			"AppVersion   := %s\n"+
			"ServerHost  := %s\n"+
			"ServerPort   := %s\n"+
			"Log File     := %s\n"+
			"App Config   := %+v\n"+
			this.AppName,
		this.AppVersion,
		this.ServerHost,
		this.ServerPort,
		this.LogConfFile,
		this.ApplicationConfig,
	)
	return s
}

// GlobalAppConfig is applicationconfig Singleton
var GlobalAppConfig = new(AppConfig)
var configEnvUpdateMap map[string]string
var globalEnvUpdateMap map[string]string
