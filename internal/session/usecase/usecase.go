package usecase

import (
	"context"
	"fmt"

	session "github.com/aclgo/grpc-jwt/internal/session"
	"github.com/aclgo/grpc-jwt/internal/session/models"
	sessionRepo "github.com/aclgo/grpc-jwt/internal/session/repository"
	sessionToken "github.com/aclgo/grpc-jwt/internal/session/token"
	"github.com/aclgo/grpc-jwt/pkg/logger"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

type sessionUC struct {
	logger      logger.Logger
	tokenRepo   session.TokenRepo
	tokenAction session.TokenAction
}

func NewSessionUC(logger logger.Logger, redisClient *redis.Client,
	secretKey string) *sessionUC {

	tokenRepo := sessionRepo.NewjwtStore(redisClient)
	tokenAction := sessionToken.NewtokenAction(secretKey)

	return &sessionUC{
		logger:      logger,
		tokenRepo:   tokenRepo,
		tokenAction: tokenAction,
	}
}

func (s *sessionUC) CreateTokens(ctx context.Context, userID, role string) (*models.Token, error) {
	return s.createTokens(ctx, userID, role)
}

func (s *sessionUC) RefreshToken(ctx context.Context, accessTTK, refreshTTK string) (*models.Token, error) {

	err := s.verifyRevogedToken(ctx, refreshTTK)

	if err == nil {
		return nil, session.ErrTokenRevoged
	}

	if err != nil && err != redis.Nil {
		return nil, err
	}

	parsedAccess, err := s.tokenAction.ParseToken(accessTTK)
	if err != nil {
		s.logger.Errorf("RefreshToken.ParseToken: %v", err)
		return nil, fmt.Errorf("RefreshToken.ParseToken: %v", err)
	}

	parsedRefresh, err := s.tokenAction.ParseToken(refreshTTK)
	if err != nil {
		s.logger.Errorf("RefreshToken.ParseToken: %v", err)
		return nil, fmt.Errorf("RefreshToken.ParseToken: %v", err)
	}

	claimsAccess, err := s.tokenAction.GetClaims(parsedAccess)
	if err != nil {
		s.logger.Errorf("RefreshToken.GetClaims: %v", err)
		return nil, fmt.Errorf("RefreshToken.GetClaims: %v", err)
	}

	claimsRefresh, err := s.tokenAction.GetClaims(parsedRefresh)
	if err != nil {
		s.logger.Errorf("RefreshToken.GetClaims: %v", err)
		return nil, fmt.Errorf("RefreshToken.GetClaims: %v", err)
	}

	idAccess, _ := claimsAccess["id"].(string)
	idRefresh, _ := claimsRefresh["id"].(string)
	typeRefresh, _ := claimsRefresh["type"].(string)

	if idAccess != idRefresh {
		return nil, session.ErrMistachTokenID
	}

	if typeRefresh != session.TypeRefreshTTK {
		return nil, session.ErrTypeTokenInvalid
	}

	role, _ := claimsAccess["role"].(string)

	return s.createTokens(ctx, idRefresh, role)
}

func (s *sessionUC) ValidToken(ctx context.Context, ttkString string) (jwt.MapClaims, error) {
	parsedAccess, err := s.tokenAction.ParseToken(ttkString)
	if err != nil {
		s.logger.Error(session.ErrInvalidToken)
		return nil, session.ErrInvalidToken
	}

	claimsAccess, err := s.tokenAction.GetClaims(parsedAccess)
	if err != nil {
		s.logger.Error(session.ErrInvalidToken)
		return nil, session.ErrInvalidToken
	}

	expUnix := claimsAccess["exp"].(float64)
	if s.tokenAction.IsExpired(expUnix) {
		s.logger.Error(session.ErrTokenExpired)
		return nil, session.ErrTokenExpired
	}

	if err := s.verifyRevogedToken(ctx, ttkString); err != nil {
		if err == redis.Nil {
			return claimsAccess, nil
		}

		return nil, err
	}

	return nil, session.ErrTokenRevoged
}

func (s *sessionUC) RevogeToken(ctx context.Context, ttkAccess, ttkRefresh string) error {

	err := s.tokenRepo.Set(
		ctx,
		ttkAccess,
		session.TtlExpAccessTTK,
	)

	if err != nil {
		s.logger.Errorf("RevogeToken.Set: %v", err)
		return fmt.Errorf("RevogeToken.Set: %v", err)
	}

	err = s.tokenRepo.Set(ctx,
		ttkRefresh,
		session.TtlExpRefreshTTK,
	)

	if err != nil {
		s.logger.Errorf("RevogeToken.Set: %v", err)
		return fmt.Errorf("RevogeToken.Set: %v", err)
	}

	return nil
}

func (s *sessionUC) VerifyRevogedToken(ctx context.Context, ttkString string) error {
	return s.verifyRevogedToken(ctx, session.FormatKeyRevogedToken(session.DefaultKeyRevogedTokenAccess, ttkString))
}

func (s *sessionUC) verifyRevogedToken(ctx context.Context, ttkString string) error {
	err := s.tokenRepo.Get(ctx, ttkString)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionUC) createTokens(ctx context.Context, userID, role string) (*models.Token, error) {
	access, err := s.tokenAction.NewToken(session.TypeAccessTTK, userID, role, session.TtlExpAccessTTK)
	if err != nil {
		return nil, err
	}

	refresh, err := s.tokenAction.NewToken(session.TypeRefreshTTK, userID, "", session.TtlExpRefreshTTK)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Access:  access,
		Refresh: refresh,
	}, nil
}
