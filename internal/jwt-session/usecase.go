package session

import (
	"context"

	"github.com/aclgo/grpc-jwt/internal/jwt-session/models"
	"github.com/golang-jwt/jwt"
)

type SessionUC interface {
	CreateTokens(context.Context) (*models.Token, error)
	RefreshToken(context.Context, string) (*models.Token, error)
	ValidToken(context.Context, string) (*jwt.MapClaims, error)
	RevogeToken(context.Context) error
	VerifyRevogedToken(context.Context) error
}
