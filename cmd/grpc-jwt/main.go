package main

import (
	"fmt"

	"github.com/aclgo/grpc-jwt/config"
	"github.com/aclgo/grpc-jwt/internal/server"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/aclgo/grpc-jwt/pkg/postgres"
	rredis "github.com/aclgo/grpc-jwt/pkg/redis"
)

func main() {
	fmt.Println("grpc-jwt")

	cfg := config.Config{}
	logger := logger.NewLogger(&cfg)

	db, err := postgres.Connect(&cfg)
	if err != nil {
		logger.Fatal(err)
	}

	redisClient := rredis.NewRedisClient(&cfg)

	server := server.NewServer(db, redisClient, logger, &cfg)

	logger.Fatal(server.Run())

}
