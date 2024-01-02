package session

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	DefaultKeyRevogedTokenAccess  = "access:"
	DefaultKeyRevogedTokenRefresh = "refresh:"
)

func FormatKeyRevogedToken(tmpl, s string) string {
	return fmt.Sprintf("%s: %s", tmpl, s)
}

type TokenAction interface {
	NewToken(typeToken, userID, role string, ttl time.Duration) (string, error)
	ParseToken(tokenString string) (*jwt.Token, error)
	GetClaims(token *jwt.Token) (jwt.MapClaims, error)
	IsExpired(timeUnix float64) bool
}
