package service

import (
	"context"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
	"github.com/google/uuid"
)

type refreshTokenService struct {
	authorRepository domain.AuthorRepository
	contextTimeout   time.Duration
}

func NewRefreshTokenService(authorRepository domain.AuthorRepository, timeout time.Duration) domain.RefreshTokenService {
	return &refreshTokenService{
		authorRepository: authorRepository,
		contextTimeout:   timeout,
	}
}

// CreateAccessToken implements domain.RefreshTokenService.
func (rts *refreshTokenService) CreateAccessToken(author *domain.Author, expired int) (accessToken string, err error) {
	return utils.CreateAccessToken(author, expired)
}

// CreateRefreshToken implements domain.RefreshTokenService.
func (rts *refreshTokenService) CreateRefreshToken(author *domain.Author, expired int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(author, expired)
}

// ExtractIDFromToken implements domain.RefreshTokenService.
func (rts *refreshTokenService) ExtractIDFromToken(tokenRequested string) (uuid.UUID, error) {
	return utils.ExtractIDFromToken(tokenRequested)
}

// GetAuthorByID implements domain.RefreshTokenService.
func (rts *refreshTokenService) GetAuthorByID(ctx context.Context, id uuid.UUID) (*domain.Author, error) {
	c, cancel := context.WithTimeout(ctx, rts.contextTimeout)
	defer cancel()
	return rts.authorRepository.GetByID(c, id)
}
