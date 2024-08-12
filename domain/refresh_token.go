package domain

import "context"

type RefreshToken struct {
	RefreshToken string `binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenService interface {
	GetAuthorByID(ctx context.Context, id string) (Author, error)
	CreateAccessToken(author *Author, expired int) (accessToken string, err error)
	CreateRefreshToken(author *Author, expired int) (refreshToken string, err error)
	ExtractIDFromToken(tokenRequested string) (string, error)
}
