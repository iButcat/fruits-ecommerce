package config

import "github.com/spf13/viper"

// Containing different fields as DSN, port, etc...
type Config struct {
	Dsn       string `mapstructure:"dsn"`
	Port      string `mapstructure:"port"`
	SecretKey string `mapstructure:"secret_key"`
}

// Load our config in our struct to use it.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.AddConfigPath(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
