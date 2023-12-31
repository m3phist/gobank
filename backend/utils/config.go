package utils

import "github.com/spf13/viper"

type Config struct {
	Port           int    `mapstructure:"PORT"`
	DB_driver      string `mapstructure:"DB_DRIVER"`
	DB_source      string `mapstructure:"DB_SOURCE"`
	DB_source_live string `mapstructure:"DB_SOURCE_LIVE"`
	Signing_key    string `mapstructure:"SIGNING_KEY"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil

}
