package domain

import "context"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginService interface {
	GetUserByUsername(ctx context.Context, username string) (*Author, error)
	CreateAccessToken(author *Author, expired int) (accessToken string, err error)
	CreateRefreshToken(author *Author, expired int) (refreshToken string, err error)
}
