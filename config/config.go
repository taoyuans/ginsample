package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ConfigValue *Config

// 全局配置
type Config struct {
	Service    string `json:"service"`
	HttpPort   string `json:"httpPort"`
	UseJwtAuth bool   `json:"useJwtAuth"`
	Mysql      struct {
		DataSourceName  string `json:"dataSourceName"`
		MaxOpenConn     int    `json:"maxOpenConn"`
		MaxIdleConn     int    `json:"maxIdleConn"`
		MaxConnLifeTime int    `json:"maxConnLifeTime"`
	} `json:"mysql"`
	Redis struct {
		DataSourceName string `json:"dataSourceName"`
		DB             int    `json:"dB"`
		Password       string `json:"password"`
	} `json:"redis"`
	Log struct {
		LogLevel      int8   `json:"logLevel"`
		LogFilePath   string `json:"logFilePath"`
		LogFileName   string `json:"logFileName"`
		LogMaxSize    int    `json:"logMaxSize"`
		LogMaxBackups int    `json:"logMaxBackups"`
		LogMaxAge     int    `json:"logMaxAge"`
	} `json:"log"`
	Server struct {
		Auth string `json:"auth"`
	}
}

func LoadConfig() error {
	var configFilePath = "config/"
	viper.SetConfigName("config")
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if err := viper.Unmarshal(&ConfigValue); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	return nil
}
