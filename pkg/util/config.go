package util

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	config Config
	once   sync.Once
)

type Config struct {
	Port                    int32  `mapstructure:"PORT"`
	UserMicroserviceDomain  string `mapstructure:"USER_MS_DOMAIN"`
	AuthMicroserviceDomain  string `mapstructure:"AUTH_MS_DOMAIN"`
	AuthMicroserviceRestUrl string `mapstructure:"AUTH_MS_REST_URL"`
}

func LoadConfig(path string) (Config, error) {
	var err error
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
	})
	return config, err
}

func GetConfig() Config {
	return config
}
