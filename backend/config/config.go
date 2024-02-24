package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	log *log.Logger
}

func GetConfig(log *log.Logger) *Config {

	c := Config{
		log: log,
	}
	c.init()
	return &c
}

func (c *Config) init() {
	c.log.Info("Initializing Config")
	viper.SetConfigName("servocoolant")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func (c *Config) GetVersion() string {
	return viper.GetString("version")
}
