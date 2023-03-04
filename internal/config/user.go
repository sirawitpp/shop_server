package config

import (
	"time"

	"github.com/spf13/viper"
)

type UserConfig struct {
	DSN               string        `mapstructure:"DSN"`
	HttpServerAddress string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	Sign              string        `mapstructure:"SIGN"`
	TokenDuration     time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadUserConfig(path string) (config UserConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("user")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
