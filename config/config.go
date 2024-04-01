package config

import (
	"context"
	"github.com/spf13/viper"
	"go-boilerplate/pkg/logger"
)

type Config struct {
	App           App
	SocialAuth    SocialAuth
	Email         Email
	Database      Database
	Redis         Redis
	Uuid          Uuid
	Otel          Otel
	Vault         VaultConfig
	CoinMarketCap CoinMarketCap
	Indexer       Indexer
	Sms           Sms
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error(context.Background()).Err(err)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		config = nil
		return config, err
	}

	return config, nil
}
