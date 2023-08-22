package config

import "time"

type Config struct {
	ApiVersion  string
	SecretKey   string
	LogLevel    string
	LogEncoding string
	Server
	Database
	Redis
}

type Server struct {
	AppVersion string `mapstructure:"APP_VERSION"`
	Mode       string `mapstructure:"SERVER_MODE"`
	Port       string `mapstructure:"SERVER_PORT"`
}

type Database struct {
	Url    string `mapstructure:"DATABASE_URL"`
	Driver string
}

type Redis struct {
	Addr string `mapstructure:"REDIS_ADDR"`
	DB   int    `mapstructure:"SERVER_PORT"`
	Pass string
}

type Jwt struct {
	ExpirateToken time.Duration
}
