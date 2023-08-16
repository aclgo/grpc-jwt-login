package config

type Config struct {
	Server
	Database
	Redis
}

type Server struct {
	Mode string
	Port string
}

type Database struct {
	Url string
}

type Redis struct {
	Url  string
	DB   int
	Pass string
}
