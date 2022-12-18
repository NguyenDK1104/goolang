package utils

import "github.com/spf13/viper"

type Config struct {
	DBUserName    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBAddress     string `mapstructure:"DB_ADDRESS"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBServiceName string `mapstructure:"DB_SERVICE_NAME"`
}

func LoadConfig(path string) (config *Config, err error) {
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
