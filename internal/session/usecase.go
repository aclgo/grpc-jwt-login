package session

import (
	"context"

	"github.com/aclgo/grpc-jwt/internal/session/models"
	"github.com/golang-jwt/jwt"
)

type SessionUC interface {
	CreateTokens(context.Context, string, string) (*models.Token, error)
	RefreshToken(context.Context, string, string) (*models.Token, error)
	ValidToken(context.Context, string) (jwt.MapClaims, error)
	RevogeToken(context.Context, string, string) error
	VerifyRevogedToken(context.Context, string) error
}
