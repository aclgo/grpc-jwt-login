package config

import "github.com/spf13/viper"

func Load(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("gprc-jwt")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
