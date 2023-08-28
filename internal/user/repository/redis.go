package repository

import "github.com/go-redis/redis"

type redisRepo struct {
	redisClient *redis.Client
}
