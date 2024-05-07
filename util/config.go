package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	DBConnectionSource   string        `mapstructure:"DB_CONNECTION_SOURCE"`
	DBMigrationSource    string        `mapstructure:"DB_MIGRATION_SOURCE"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EncryptionKey        string        `mapstructure:"KEY"`
}

func LoadConfig(path string) (Config, error) {

	config := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, err
}
