package rredis

import (
	"github.com/aclgo/grpc-jwt/config"
	"github.com/go-redis/redis"
)

func NewRedisClient(c *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		DB:       c.RedisDB,
		Password: c.RedisPass,
		PoolSize: 100,
	})

	return client
}
