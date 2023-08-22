package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type jwtRepo struct {
	redisClient *redis.Client
}

func NewjwtStore(redisClient *redis.Client) *jwtRepo {
	return &jwtRepo{
		redisClient: redisClient,
	}
}

func (j *jwtRepo) Get(ctx context.Context, tokenString string) error {
	_, err := j.redisClient.Get(tokenString).Result()
	return err
}

func (j *jwtRepo) Set(ctx context.Context, tokenString string, ttl time.Duration) error {
	return j.redisClient.Set(tokenString, nil, ttl).Err()
}
