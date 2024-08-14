package domain

import (
	"context"

	"github.com/google/uuid"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}



type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenService interface {
	GetAuthorByID(ctx context.Context, id uuid.UUID) (*Author, error)
	CreateAccessToken(author *Author, expired int) (accessToken string, err error)
	CreateRefreshToken(author *Author, expired int) (refreshToken string, err error)
	ExtractIDFromToken(tokenRequested string) (uuid.UUID, error)
}
