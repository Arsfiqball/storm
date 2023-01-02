package provider

import (
	"github.com/spf13/viper"
)

func ProvideViper() *viper.Viper {
	config := viper.New()

	config.SetDefault("ADDRESS", ":3000")

	config.SetEnvPrefix("APP")
	config.AllowEmptyEnv(true)
	config.AutomaticEnv()

	return config
}
