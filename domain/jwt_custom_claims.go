package domain

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
