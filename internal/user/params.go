package user

import (
	"context"
	"time"

	"github.com/aclgo/grpc-jwt/internal/models"
)

type ParamsCreateUser struct {
	Name     string
	Lastname string
	Password string
	Email    string
	Role     string
}

func (p *ParamsCreateUser) Validate(ctx context.Context) error {
	return nil
}

type ParamsUpdateUser struct{}

func (p *ParamsUpdateUser) Validate(ctx context.Context) error {
	return nil
}

type ParamsOutputUser struct {
	Id        string
	Name      string
	Lastname  string
	Password  string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Dto(user *models.User) *ParamsOutputUser {
	return &ParamsOutputUser{
		Id:        user.Id,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Password:  user.Password,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
