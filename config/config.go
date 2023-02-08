package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ConfigValue *Config

type Config struct {
	GinMode  string
	HttpPort string
	Service  string
	DataBase struct {
		DriverName      string
		DataSourceName  string
		MaxOpenConn     string
		MaxIdleConn     string
		MaxConnLifeTime string
	}
	Log struct {
		Useable     bool
		LogFilePath string
		LogFileName string
	}
}

func Read(env string, config *Config) error {
	var configFilePath = "config/"
	if env == "" {
		env = "dev"
	} else if env == "test" {
		configFilePath = "../../config/"
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if env != "" {
		f, err := os.Open(configFilePath + env + "/config.yml")
		if err != nil {
			return fmt.Errorf("Fatal error config file: %s \n", err)
		}
		defer f.Close()
		viper.MergeConfig(f)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	ConfigValue = config
	return nil
}
