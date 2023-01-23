package config

import "github.com/spf13/viper"

type EnvVars struct {
	PORT    string `mapstructure:"PORT"`
	JWT_KEY string `mapstructure:"JWT_KEY"`
}

func LoadConfig() (config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// TODO: add validation

	err = viper.Unmarshal(&config)
	return
}
