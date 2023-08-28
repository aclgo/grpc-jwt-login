package grpc_errors

import (
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

func ParseGRPCErrors(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, redis.Nil):
		return codes.NotFound
	case errors.Is(err, sql.ErrNoRows):

	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, sql.ErrNoRows):

	}

	return codes.Internal
}
