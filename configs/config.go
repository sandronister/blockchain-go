package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBPath string `mapstructure:"DB_PATH"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("blockchain-go")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, using default values")
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
