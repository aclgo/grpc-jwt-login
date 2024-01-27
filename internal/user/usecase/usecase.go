package usecase

import (
	"context"

	"fmt"
	"time"

	"github.com/aclgo/grpc-jwt/internal/models"
	session "github.com/aclgo/grpc-jwt/internal/session"
	"github.com/aclgo/grpc-jwt/internal/user"
	"github.com/aclgo/grpc-jwt/internal/utils"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	ClientRole = "client"
)

var (
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrEmailCadastred    = errors.New("email cadastred")
	ErrInvalidEmail      = errors.New("email invalid")
	ErrPasswordSmall     = errors.New("password small lenght")
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
	if !utils.ValidMail(params.Email) {
		u.logger.Errorf("Register.FindByEmail: %v", ErrInvalidEmail)
		return nil, fmt.Errorf("Register.FindByEmail: %v", ErrInvalidEmail)
	}

	foundUser, _ := u.userRepoDatabase.FindByEmail(ctx, params.Email)
	if foundUser != nil {
		u.logger.Errorf("Register.FindByEmail: %v", ErrEmailCadastred)
		return nil, fmt.Errorf("Register.FindByEmail: %v", ErrEmailCadastred)
	}

	created, err := u.userRepoDatabase.Add(ctx, &models.User{
		UserID:    uuid.NewString(),
		Name:      params.Name,
		Lastname:  params.Lastname,
		Password:  params.HashPass(),
		Email:     params.Email,
		Role:      ClientRole,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		u.logger.Errorf("Register.Add: %v", err)
		return nil, fmt.Errorf("Register.Add: %v", err)
	}

	return user.Dto(created), nil
}

func (u *userUC) Login(ctx context.Context, email string, password string) (*models.Tokens, error) {
	// fmt.Println("init login")
	foundUser, err := u.userRepoDatabase.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("Login.FindByEmail: %v", err)
		return nil, fmt.Errorf("Login.FindByEmail: %v", err)
	}

	// fmt.Println("found user")

	if err := foundUser.ComparePass(password); err != nil {
		u.logger.Errorf("Login: %v", ErrPasswordIncorrect)
		return nil, ErrPasswordIncorrect
	}

	// fmt.Println("compare pass")

	if err := u.userRepoCache.Set(ctx, foundUser); err != nil {
		u.logger.Warn("Login.Set: %v", err)
	}

	tokens, err := u.jwtSession.CreateTokens(ctx, foundUser.UserID, foundUser.Role)
	if err != nil {
		u.logger.Errorf("Login.CreateTokens: %v", err)
		return nil, fmt.Errorf("Login.CreateTokens: %v", err)
	}

	// u.logger.Info(tokens.Access)
	// u.logger.Info(tokens.Refresh)

	return &models.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}, nil

}

func (u *userUC) Logout(ctx context.Context, accessTTK string, refreshTTK string) error {
	err := u.jwtSession.RevogeToken(ctx, accessTTK, refreshTTK)
	if err != nil {
		u.logger.Errorf("Logout.RevogeToken: %v", err)
		return fmt.Errorf("Logout.RevogeToken: %v", err)
	}

	return nil
}

func (u *userUC) FindByID(ctx context.Context, userID string) (*user.ParamsOutputUser, error) {
	var (
		foundUser *models.User
		err       error
	)

	foundUser, err = u.userRepoCache.Get(ctx, userID)
	if err == redis.Nil {
		foundUser, err = u.userRepoDatabase.FindByID(ctx, userID)
		if err != nil {
			u.logger.Errorf("FindByID: %v", err)
			return nil, fmt.Errorf("FindByID: %v", err)
		}

		if err := u.userRepoCache.Set(ctx, foundUser); err != nil {
			u.logger.Warn("FindByID.Set: %v", err)
		}

		return user.Dto(foundUser), nil
	}

	if err != nil {
		u.logger.Errorf("FindByID.Get: %v", err)
		return nil, fmt.Errorf("FindByID.Get: %v", err)
	}

	return user.Dto(foundUser), nil
}

func (u *userUC) FindByEmail(ctx context.Context, userEmail string) (*user.ParamsOutputUser, error) {

	var (
		foundUser *models.User
		err       error
	)

	foundUser, err = u.userRepoCache.Get(ctx, userEmail)
	if err == redis.Nil {
		foundUser, err = u.userRepoDatabase.FindByEmail(ctx, userEmail)
		if err != nil {
			u.logger.Errorf("FindByEmail: %v", err)
			return nil, fmt.Errorf("FindByEmail: %v", err)
		}

		if err := u.userRepoCache.Set(ctx, foundUser); err != nil {
			u.logger.Warn("FindByEmail.Set: %v", err)
		}

		return user.Dto(foundUser), nil
	}

	if err != nil {
		u.logger.Errorf("FindByEmail.Get: %v", err)
		return nil, fmt.Errorf("FindByEmail.Get: %v", err)
	}

	return user.Dto(foundUser), nil
}

func (u *userUC) Update(ctx context.Context, params *user.ParamsUpdateUser) (*user.ParamsOutputUser, error) {
	if err := params.Validate(ctx); err != nil {
		u.logger.Errorf("Update.Validate: %v", err)
		return nil, errors.Wrap(err, "Update.Validate")
	}

	newUser, err := u.userRepoDatabase.Update(ctx,
		&models.User{
			UserID:    params.UserID,
			Name:      params.Name,
			Lastname:  params.Lastname,
			Password:  params.Password,
			Email:     params.Email,
			Verified:  params.Verified,
			UpdatedAt: time.Now(),
		},
	)

	if err != nil {
		u.logger.Errorf("Update.Update: %v", err)
		return nil, errors.Wrap(err, "Update.Update")
	}

	return user.Dto(newUser), nil
}

func (u *userUC) ValidToken(ctx context.Context, params *user.ParamsValidToken) (*user.ParamsJwtData, error) {
	claims, err := u.jwtSession.ValidToken(ctx, params.AccessToken)
	if err != nil && !errors.Is(err, redis.Nil) {
		u.logger.Errorf("ValidToken: %v", err)
		return nil, err
	}
	// u.logger.Info(claims)

	return &user.ParamsJwtData{
		UserID: claims["id"].(string),
		Role:   claims["role"].(string),
	}, nil
}

func (u *userUC) RefreshTokens(ctx context.Context, params *user.ParamsRefreshTokens) (*user.RefreshTokens, error) {
	tokens, err := u.jwtSession.RefreshToken(ctx, params.AccessToken, params.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &user.RefreshTokens{
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	}, nil
}
