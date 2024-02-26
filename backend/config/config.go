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

type Tool struct {
	Length float32 `json:"length" mapstructure:"length"` //must be capitalized in struct for viper to unmarshall
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
	
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	c.Viper.WatchConfig()

}

func (c *Config) GetVersion() string {
	return c.Viper.GetString("version")
}

func (c *Config) GetAllToolLengths() (map[string]string, error) {

	// var tools map[string]Tool
	// err := c.Viper.UnmarshalKey("tools", &tools)

	tools := c.Viper.GetStringMapString("tools")

	// if err != nil {

	// 	c.log.Error(err.Error())
	// 	return nil, err
	// }
	return tools, nil
}

func (c *Config) GetToolLength(toolNumber int) (*string, error) {

	tools, err := c.GetAllToolLengths()

	if err != nil {
		return nil, err
	}

	tool, ok := tools[fmt.Sprint(toolNumber)]
	if ok {
		return &tool, nil
	}
	err = fmt.Errorf("tool %v not found", toolNumber)
	c.log.Error(err)
	return nil, err
}

func (c *Config) SetToolLength(toolNumber int, length float32) {

	c.Viper.Set("tools.12", fmt.Sprint(length))
	// c.Viper.Set(fmt.Sprintf("tools.%v.length", toolNumber), length)
	err := c.Viper.WriteConfig()

	if err != nil {
		c.log.Error(fmt.Sprintf("error writing config %v", err))
	}
	// err = c.Viper.ReadInConfig()
	// if err != nil {
	// 	c.log.Error(fmt.Sprintf("error reading in config %v", err))
	// }

}
