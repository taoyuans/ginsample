package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ConfigValue *Config

type Config struct {
	GinMode  string
	DataBase struct {
		DriverName      string
		DataSourceName  string
		MaxOpenConn     string
		MaxIdleConn     string
		MaxConnLifeTime string
	}
	HttpPort string
	Service  string
}

func Read(env string, config *Config) error {
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("config/.")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if env != "" {
		f, err := os.Open("config/" + env + "/config.yml")
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
