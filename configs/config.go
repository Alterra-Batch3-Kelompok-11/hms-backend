package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
	ApiPort    string
	JwtKey     string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	err := viper.Unmarshal(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	Cfg = cfg
}
