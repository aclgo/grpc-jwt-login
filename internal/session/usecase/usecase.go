package usecase

import (
	"context"
	"fmt"
	"time"

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

const (
	typeAccessTTK    = "access"
	typeRefreshTTK   = "refresh"
	ttlExpAccessTTK  = time.Hour
	ttlExpRefreshTTK = time.Hour * 24
)

func (s *sessionUC) CreateTokens(ctx context.Context, userID, role string) (*models.Token, error) {
	return s.createTokens(ctx, userID, role)
}

func (s *sessionUC) RefreshToken(ctx context.Context, accessTTK, refreshTTK string) (*models.Token, error) {

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

	// idAccess, _ := claimsAccess["id"].(string)
	idRefresh, _ := claimsRefresh["id"].(string)
	typeTTK, _ := claimsRefresh["type"].(string)

	if typeTTK != typeRefreshTTK {
		return nil, ErrTypeTokenInvalid
	}

	role, _ := claimsAccess["role"].(string)

	return s.createTokens(ctx, idRefresh, role)
}

func (s *sessionUC) ValidToken(ctx context.Context, ttkString string) (jwt.MapClaims, error) {
	parsedAccess, err := s.tokenAction.ParseToken(ttkString)
	if err != nil {
		s.logger.Error(ErrInvalidToken)
		return nil, ErrInvalidToken
	}

	claimsAccess, err := s.tokenAction.GetClaims(parsedAccess)
	if err != nil {
		s.logger.Error(ErrInvalidToken)
		return nil, ErrInvalidToken
	}

	expUnix := claimsAccess["exp"].(float64)
	if s.tokenAction.IsExpired(expUnix) {
		s.logger.Error(ErrTokenExpired)
		return nil, ErrTokenExpired
	}

	if !s.tokenAction.IsExpired(expUnix) {
		return claimsAccess, nil
	}

	s.logger.Error(ErrTokenRevoged)
	return nil, ErrTokenRevoged
}

func (s *sessionUC) RevogeToken(ctx context.Context, ttkAccess, ttkRefresh string) error {
	token, err := s.tokenAction.ParseToken(ttkAccess)
	if err != nil {
		return fmt.Errorf("RevogeToken.ParseToken: %v", err)
	}

	parsedTTK, err := s.tokenAction.GetClaims(token)
	if err != nil {
		s.logger.Errorf("RevogeToken.GetClaims: %v", err)
		return fmt.Errorf("RevogeToken.GetClaims: %v", err)
	}

	exp := parsedTTK["exp"].(float64)
	now := time.Now().Unix()

	timeRestant := exp - float64(now)

	if s.tokenRepo.Set(ctx, ttkAccess, time.Duration(timeRestant)*time.Second); err != nil {
		s.logger.Errorf("RevogeToken.Set: %v", err)
		return fmt.Errorf("RevogeToken.Set: %v", err)
	}

	return nil
}

func (s *sessionUC) VerifyRevogedToken(ctx context.Context, ttkString string) error {
	err := s.tokenRepo.Get(ctx, ttkString)
	if err == redis.Nil {
		return nil
	}
	s.logger.Error(ErrTokenRevoged)
	return ErrTokenRevoged
}

func (s *sessionUC) createTokens(ctx context.Context, userID, role string) (*models.Token, error) {
	access, err := s.tokenAction.NewToken(typeAccessTTK, userID, role, ttlExpAccessTTK)
	if err != nil {
		return nil, err
	}

	refresh, err := s.tokenAction.NewToken(typeRefreshTTK, userID, "", ttlExpRefreshTTK)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Access:  access,
		Refresh: refresh,
	}, nil
}
