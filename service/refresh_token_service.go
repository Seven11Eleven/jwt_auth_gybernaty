package service

import (
	"context"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
)

type RefreshTokenService struct{
	authorRepository domain.AuthorRepository
	contextTimeout time.Duration
}

func NewRefreshTokenService(authorRepository domain.ArticleRepository, timeout time.Duration) domain.RefreshToken