package grpc_errors

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
	"google.golang.org/grpc/codes"
)

type EmptyCredentials struct {
}

func (e EmptyCredentials) Error() string {
	return "empty credentials"
}

var ErrNoCtxMetaData = errors.New("no ctx metadata")
var ErrInvalidToken = errors.New("invalid token")

func ParseGRPCErrors(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, redis.Nil):
		return codes.NotFound

	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	}

	return codes.Internal
}
