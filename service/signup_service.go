package service

import (
	"context"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
)

type signupService struct {
	authorRepository domain.AuthorRepository
	contextTimeout   time.Duration
}

// CheckUsernameExists implements domain.SignUpService.
func (sus *signupService) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	c, cancel := context.WithTimeout(ctx, time.Duration(sus.contextTimeout))
	defer cancel()
	exists, err := sus.authorRepository.CheckUsernameExists(c, username)
	if err != nil {
        return false, err
    }
    return exists, nil
}

// Create implements domain.SignUpService.
func (sus *signupService) Create(ctx context.Context, author *domain.Author) error {
	if !isAlpha(author.Username) {
		return domain.ErrInvalidUsername
	}
	
	c, cancel := context.WithTimeout(ctx, sus.contextTimeout)
	defer cancel()
	return sus.authorRepository.Create(c, author)
}

// CreateAccessToken implements domain.SignUpService.
func (sus *signupService) CreateAccessToken(author *domain.Author, expired int) (accessToken string, err error) {
	return utils.CreateAccessToken(author, expired)
}

// CreateRefreshToken implements domain.SignUpService.
func (sus *signupService) CreateRefreshToken(author *domain.Author, expired int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(author, expired)
}

// GetUserByUsername implements domain.SignUpService.
func (sus *signupService) GetUserByUsername(ctx context.Context, username string) (*domain.Author, error) {
	c, cancel := context.WithTimeout(ctx, sus.contextTimeout)
	defer cancel()
	return sus.authorRepository.GetByUsername(c, username)
}

func NewSignupService(authorRepository domain.AuthorRepository, timeout time.Duration) domain.SignUpService {
	return &signupService{
		authorRepository: authorRepository,
		contextTimeout:   timeout,
	}
}
