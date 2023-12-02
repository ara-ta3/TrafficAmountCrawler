package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	c, e := Load()
	fmt.Printf("%+v, %+v", c, e)
}


func Load() (config *EnvConfigs, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}

type EnvConfigs struct {
	NihonTsushinID   string `mapstructure:"NIHON_TSUSHIN_ID"`
	NihonTsushinPass string `mapstructure:"NIHON_TSUSHIN_PASS"`
}

