package user

import (
	"context"
	"errors"
	"time"

	"github.com/aclgo/grpc-jwt/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type ParamsCreateUser struct {
	Name     string
	Lastname string
	Password string
	Email    string
}

func (p *ParamsCreateUser) HashPass() string {
	bc, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	return string(bc)
}

func (p *ParamsCreateUser) Validate(ctx context.Context) error {
	return nil
}

type ParamsUpdateUser struct {
	UserID    string
	Name      string
	Lastname  string
	Password  string
	Email     string
	UpdatedAt time.Time
}

func (p *ParamsUpdateUser) Validate(ctx context.Context) error {
	if p.UserID == "" {
		return errors.New("user id empty")
	}
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

func (p *ParamsOutputUser) ClearPass() {
	p.Password = ""
}

func Dto(user *models.User) *ParamsOutputUser {
	return &ParamsOutputUser{
		Id:        user.UserID,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Password:  "",
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type ParamsJwtData struct {
	UserID string
	Role   string
}

type ParamsValidToken struct {
	AccessToken string
}

type ParamsRefreshTokens struct {
	AccessToken  string
	RefreshToken string
}

type RefreshTokens struct {
	AccessToken  string
	RefreshToken string
}
