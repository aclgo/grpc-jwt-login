package usecase

import (
	"context"
	"errors"
	"time"

	session "github.com/aclgo/grpc-jwt/internal/jwt-session"
	"github.com/aclgo/grpc-jwt/internal/jwt-session/models"
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

type sessionUC struct {
	logger      logger.Logger
	tokenRepo   session.TokenRepo
	tokenAction session.TokenAction
}

func NewSessionUC(logger logger.Logger, tokenRepo session.TokenRepo,
	tokenAction session.TokenAction) *sessionUC {
	return &sessionUC{
		logger:      logger,
		tokenRepo:   tokenRepo,
		tokenAction: tokenAction,
	}
}

func (s *sessionUC) CreateTokens(context.Context) (*models.Token, error)         { return nil, nil }
func (s *sessionUC) RefreshToken(context.Context, string) (*models.Token, error) { return nil, nil }
func (s *sessionUC) ValidToken(context.Context, string) (*jwt.MapClaims, error)  { return nil, nil }

func (s *sessionUC) RevogeToken(ctx context.Context, tokenString string) error {
	parsedToken, err := s.tokenAction.ParseToken(tokenString)
	if err != nil {
		return err
	}

	claims, err := s.tokenAction.GetClaims(parsedToken)
	if err != nil {
		return err
	}

	exp := claims["exp"].(float64)
	now := time.Now().Unix()

	timeRestant := exp - float64(now)

	return s.tokenRepo.Set(ctx, tokenString, time.Duration(timeRestant))
}

func (s *sessionUC) VerifyRevogedToken(ctx context.Context, tokenString string) error {
	err := s.tokenRepo.Get(ctx, tokenString)
	if err == redis.Nil {
		return nil
	}

	return errors.New("token revoged")
}
