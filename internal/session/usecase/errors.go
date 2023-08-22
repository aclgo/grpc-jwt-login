package usecase

import "errors"

var (
	ErrTokenRevoged     = errors.New("token revoged")
	ErrTokenExpired     = errors.New("token expired")
	ErrInvalidToken     = errors.New("token invalid")
	ErrTypeTokenInvalid = errors.New("type token invalid")
)
