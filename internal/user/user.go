package user

import (
	"context"

	"github.com/aclgo/grpc-jwt/internal/models"
)

type UserRepoDatabase interface {
	Add(context.Context, *models.User) (*models.User, error)
	FindByID(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	Update(context.Context, *models.User) (*models.User, error)
}

type UserRepoCache interface {
	Set(context.Context)
	Get(context.Context)
	Del(context.Context)
}

type UserUC interface {
	Register(context.Context, *ParamsCreateUser) (*ParamsOutputUser, error)
	FindByID(context.Context, string) (*ParamsOutputUser, error)
	FindByEmail(context.Context, string) (*ParamsOutputUser, error)
	Update(context.Context, *ParamsUpdateUser) (*ParamsOutputUser, error)
	Login(context.Context, string, string) (*models.Tokens, error)
	Logout(context.Context, string, string) error
}
