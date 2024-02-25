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
	Length float32 `mapstructure:"length"` //must be capitalized in struct for viper to unmarshall
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

}

func (c *Config) GetVersion() string {
	return c.Viper.GetString("version")
}

func (c *Config) GetToolLength(toolNumber int) (*float32, error) {

	var tools map[int]tool
	err := c.Viper.UnmarshalKey("tools", &tools)

	if err != nil {

		c.log.Error(err.Error())
		return nil, err
	}

	tool, ok := tools[toolNumber]
	if ok {
		return &tool.Length, nil
	}
	err = fmt.Errorf("tool %v not found", toolNumber)
	c.log.Error(err)
	return nil, err
}

func (c *Config) SetToolLength(toolNumber int, length float32) {

	c.Viper.Set(fmt.Sprintf("tools.%v.length", toolNumber), length)
	c.Viper.WriteConfig()
}
