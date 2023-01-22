package config

import "github.com/spf13/viper"

type EnvVars struct {
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_DOMAIN"`
	PORT           string `mapstructure:"AUTH0_DOMAIN"`
}

func load(config EnvVars, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// validation

	err = viper.Unmarshal(&config)
	return
}
