package config

import "github.com/spf13/viper"

type EnvVars struct {
	PORT   string `mapstructure:"PORT"`
	DB_DSN string `mapstructure:"DB_DSN"`

	JWT_SECRET     string `mapstructure:"JWT_SECRET"`
	JWT_ISSUER     string `mapstructure:"JWT_ISSUER"`
	JWT_COOKIE_KEY string `mapstructure:"JWT_COOKIE_KEY"`
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
