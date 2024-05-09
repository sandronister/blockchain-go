package config

import "github.com/spf13/viper"

type Config struct {
	DBPath string `mapstructure:"DB_PATH"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	viper.SetConfigName("blockchain-go")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
