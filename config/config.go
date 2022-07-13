package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ExecuteMode string

type Config struct {
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

func Read(env string, config interface{}) error {
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
	return nil
}
