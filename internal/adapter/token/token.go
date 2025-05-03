package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken      = errors.New("token is invalid")
	ErrSecretKeyNotFound = errors.New("secret key not found")
	ErrAccessTokenExpired = errors.New("access token expired")
)

var _ Token = (*token)(nil)

type Token interface {
	// ValidateToken - валидирует токен и возвращает его содержимое
	ValidateToken(tokenString string) (bool, uuid.UUID, error)
}

type token struct {
	secretKey string
}

func New(secretKey string) (*token, error) {
	if secretKey == "" {
		return nil, ErrSecretKeyNotFound
	}
	return &token{
		secretKey: secretKey,
	}, nil
}

func (t *token) ValidateToken(tokenString string) (bool, uuid.UUID, error) {
	claims, err := t.parseToken(tokenString)
	if err != nil {
		return false, uuid.Nil, err
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.IsZero() {
		return false, uuid.Nil, ErrAccessTokenExpired
	}
	userId := uuid.MustParse(claims.Subject)

	return true, userId, nil
}

func (t *token) parseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(j *jwt.Token) (any, error) {
		return []byte(t.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok && token.Valid {
		return nil, ErrInvalidToken
	}
	if claims.Subject == "" {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
