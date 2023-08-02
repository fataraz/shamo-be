package config

import (
	"fmt"
	"github.com/spf13/viper"
	"shamo-be/internal/shared/database"
)

// Config ...
type Config struct {
	Apps     Apps `json:"apps"`
	Database DB   `json:"database"`
}

// Apps ...
type Apps struct {
	Name     string `json:"name"`
	HttpPort int    `json:"http_port"`
	Version  string `json:"version"`
}

// DB ...
type DB struct {
	Db database.ConfigDatabase `json:"db"`
}

func (c *Config) AppAdress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}

// NewConfig ...
func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
