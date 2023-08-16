package rredis

import (
	"github.com/aclgo/grpc-jwt/config"
	"github.com/go-redis/redis"
)

func NewRedisClient(c *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Url,
		DB:       c.Redis.DB,
		Password: c.Redis.Pass,
		PoolSize: 100,
	})

	return client
}
