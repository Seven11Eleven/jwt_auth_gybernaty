package service

import (
	"context"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
)

type loginService struct {
	authorRepository domain.AuthorRepository
	contextTimeout   time.Duration
}

func NewLoginService(authorRepository domain.AuthorRepository, timeout time.Duration) domain.LoginService {
	return &loginService{
		authorRepository: authorRepository,
		contextTimeout:   timeout,
	}
}

// CreateAccessToken implements domain.LoginService.
func (ls *loginService) CreateAccessToken(author *domain.Author, expired int) (accessToken string, err error) {
	return utils.CreateAccessToken(author, expired)
}

// CreateRefreshToken implements domain.LoginService.
func (ls *loginService) CreateRefreshToken(author *domain.Author, expired int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(author, expired)
}

// GetUserByUsername implements domain.LoginService.
func (ls *loginService) GetUserByUsername(ctx context.Context, username string) (*domain.Author, error) {
	c, cancel := context.WithTimeout(ctx, ls.contextTimeout)
	defer cancel()
	return ls.authorRepository.GetByUsername(c, username)
}
