package service

import (
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/aclgo/grpc-jwt/proto"
)

type UserService struct {
	logger logger.Logger
	userUC user.UserUC
	proto.UnimplementedUserServiceServer
}

func NewUserService(logger logger.Logger, userUC user.UserUC) *UserService {
	return &UserService{
		logger: logger,
		userUC: userUC,
	}
}
