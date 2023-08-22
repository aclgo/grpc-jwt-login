package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aclgo/grpc-jwt/internal/models"
	session "github.com/aclgo/grpc-jwt/internal/session"
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/google/uuid"
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

func (u *userUC) Register(ctx context.Context, params *user.ParamsCreateUser) (*user.ParamsOutputUser, error) {
	foundUser, err := u.userRepoDatabase.FindByEmail(ctx, params.Email)
	if foundUser != nil {
		return nil, errors.New("Register.FindByEmail: email cadastred")
	}

	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("Register.FindByEmail: %v", err)
	}

	created, err := u.userRepoDatabase.Add(ctx, &models.User{
		Id:        uuid.NewString(),
		Name:      params.Name,
		Lastname:  params.Lastname,
		Password:  params.Password,
		Email:     params.Email,
		Role:      params.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		// u.logger.Errorf("u.userRepoDatabase.Add: %v", err)
		return nil, fmt.Errorf("u.userRepoDatabase.Add: %v", err)
	}

	return user.Dto(created), nil
}

func (u *userUC) Login(ctx context.Context, email string, password string) (*models.Tokens, error) {
	foundUser, err := u.userRepoDatabase.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("Login.FindByEmail: %v", err)
	}

	if err := foundUser.ComparePass(password); err != nil {
		return nil, errors.New("password incorrect")
	}

	tokens, err := u.jwtSession.CreateTokens(ctx, foundUser.Id, foundUser.Role)
	if err != nil {
		return nil, fmt.Errorf("Login.CreateTokens: %v", err)
	}

	return &models.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}, nil
}

func (u *userUC) Logout(ctx context.Context, accessTTK string, refreshTTK string) error {
	err := u.jwtSession.RevogeToken(ctx, accessTTK, refreshTTK)
	if err != nil {
		return fmt.Errorf("Logout.RevogeToken: %v", err)
	}

	return nil
}

func (u *userUC) FindByID(ctx context.Context, userID string) (*user.ParamsOutputUser, error) {
	foundUser, err := u.userRepoDatabase.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("FindByID: %v", err)
	}

	return user.Dto(foundUser), nil
}
func (u *userUC) FindByEmail(ctx context.Context, userEmail string) (*user.ParamsOutputUser, error) {
	foundUser, err := u.userRepoDatabase.FindByEmail(ctx, userEmail)
	if err != nil {
		return nil, fmt.Errorf("FindByEmail: %v", err)
	}
	return user.Dto(foundUser), nil
}
func (u *userUC) Update(ctx context.Context, user *user.ParamsUpdateUser) (*user.ParamsOutputUser, error) {
	return nil, nil
}
