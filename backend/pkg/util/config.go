package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool
}

func LoadConfig() (Config, error) {

	output := Config{}
	viper.SetDefault("port", "50000")
	viper.SetDefault("debug", false)

	err := viper.BindEnv("debug", "DEBUG")
	if err != nil {
		return output, err
	}

	output.Debug = viper.GetBool("debug")

	return output, nil

}
