package rredis

import (
	"log"

	"github.com/aclgo/grpc-jwt/config"
	"github.com/go-redis/redis"
)

func NewRedisClient(c *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		DB:       c.RedisDB,
		Password: c.RedisPass,
		PoolSize: 10000,
	})

	if err := client.Ping().Err(); err != nil {
		log.Fatalf("NewRedisClient.Ping: %v", err)
	}

	return client
}
