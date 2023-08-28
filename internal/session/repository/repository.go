package repository

import (
	"context"
	"fmt"
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
	if err != nil {
		return fmt.Errorf("j.redisClient.Get: %v", err)
	}
	return nil
}

func (j *jwtRepo) Set(ctx context.Context, tokenString string, ttl time.Duration) error {
	err := j.redisClient.Set(tokenString, nil, ttl).Err()
	if err != nil {
		return fmt.Errorf("j.redisClient.Set: %v", err)
	}
	return nil
}
