package util

import "time"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
}
