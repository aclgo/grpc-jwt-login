package usecase

import (
	session "github.com/aclgo/grpc-jwt/internal/jwt-session"
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/pkg/logger"
)

type userUC struct {
	logger           logger.Logger
	userRepoDatabase user.UserRepoDatabase
	userRepoCache    user.UserRepoCache
	jwtSession       session.SessionUC
}

func NewUserUC(logger logger.Logger,
	userRepoDatabase user.UserRepoDatabase,
	userRepoCache user.UserRepoCache, sessionUC session.SessionUC) *userUC {
	return &userUC{
		logger:           logger,
		userRepoDatabase: userRepoDatabase,
		userRepoCache:    userRepoCache,
		jwtSession:       sessionUC,
	}
}
