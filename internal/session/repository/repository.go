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
	err := j.redisClient.Get(tokenString).Err()
	if err != nil {
		return err
	}
	return nil
}

func (j *jwtRepo) Set(ctx context.Context, tokenString string, ttl time.Duration) error {
	err := j.redisClient.Set(tokenString, nil, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
