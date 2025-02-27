package bootstrap

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	config     Config
	onceConfig sync.Once
)

type Config struct {
	Port                              int32  `mapstructure:"PORT"`
	UserMicroserviceDomain            string `mapstructure:"USER_MS_DOMAIN"`
	AuthMicroserviceDomain            string `mapstructure:"AUTH_MS_DOMAIN"`
	ClassManagementMicroserviceDomain string `mapstructure:"CLASS_MANAGEMENT_MS_DOMAIN"`
	AllowedOrigins                    string `mapstructure:"ALLOWED_ORIGINS"`
}

func LoadConfig(path string) (Config, error) {
	var err error
	onceConfig.Do(func() {
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
