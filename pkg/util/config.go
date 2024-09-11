package util

import (
	"github.com/spf13/viper"
)

// Config contains the configuration of the application.
// The values are read from environment variables.
type Config struct {
	Port                   int32  `mapstructure:"PORT"`
	UserMicroserviceDomain string `mapstructure:"USER_MS_DOMAIN"`
	AuthMicroserviceDomain string `mapstructure:"AUTH_MS_DOMAIN"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
