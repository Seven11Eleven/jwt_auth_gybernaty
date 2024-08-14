package domain

import "context"

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUpService interface {
	CheckUsernameExists(ctx context.Context, username string) (bool, error)

	Create(ctx context.Context, author *Author) error
	GetUserByUsername(ctx context.Context, username string) (*Author, error)
	CreateAccessToken(author *Author, expired int) (accessToken string, err error)
	CreateRefreshToken(author *Author, expired int) (refreshToken string, err error)
}
