package session

import (
	"context"
	"time"
)

type TokenRepo interface {
	Get(ctx context.Context, tokenString string) error
	Set(ctx context.Context, tokenString string, ttl time.Duration) error
}
