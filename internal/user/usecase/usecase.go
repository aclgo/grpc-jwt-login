package usecase

import (
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/pkg/logger"
)

type UserUC struct {
	logger    logger.Logger
	postRepo  user.UserRepoCache
	redisRepo user.UserRepoCache
}

func NewUserUC() *UserUC {
	return &UserUC{}
}
