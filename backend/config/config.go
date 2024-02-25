package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	log   *log.Logger
	Viper *viper.Viper
}

type tool struct {
	length int
}

func GetConfig(log *log.Logger) *Config {

	c := Config{
		log:   log,
		Viper: viper.GetViper(),
	}
	c.init()
	return &c
}

func (c *Config) init() {
	c.log.Info("initializing config")
	c.Viper.SetConfigName("servocoolant")
	c.Viper.SetConfigType("yaml")
	c.Viper.AddConfigPath(".")
	c.Viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var tools map[int]tool
	err = c.Viper.UnmarshalKey("tools", &tools)

	if err != nil {
		c.log.Error("error unmarshalling tools")
	}
	for count, tool := range tools {
		c.log.Info(fmt.Sprintf("found tool %v: length:%v", count, tool.length))
	}
}

func (c *Config) GetVersion() string {
	return c.Viper.GetString("version")
}
