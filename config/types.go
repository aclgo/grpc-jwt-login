package config

import "time"

type Config struct {
	ApiVersion     string        `mapstructure:"API_VERSION"`
	SecretKey      string        `mapstructure:"SECRET_KEY"`
	LogLevel       string        `mapstructure:"LOG_LEVEL"`
	LogEncoding    string        `mapstructure:"LOG_ENCODING"`
	ServerMode     string        `mapstructure:"SERVER_MODE"`
	ServerPort     string        `mapstructure:"SERVER_PORT"`
	DatabaseUrl    string        `mapstructure:"DATABASE_URL"`
	DatabaseDriver string        `mapstructure:"DRIVER_DATABASE"`
	RedisAddr      string        `mapstructure:"REDIS_ADDR"`
	RedisDB        int           `mapstructure:"REDIS_DB"`
	RedisPass      string        `mapstructure:"REDIS_PASS"`
	ExpirateToken  time.Duration `mapstructure:"EXPIRATE_TOKEN"`
}
