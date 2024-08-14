package service

import (
	"context"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/google/uuid"
)

type authorService struct {
	authorRepository domain.AuthorRepository
	contextTimeout   time.Duration
}

// CheckUsernameExists implements domain.AuthorService.
func (as *authorService) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	c, cancel := context.WithTimeout(ctx, time.Duration(as.contextTimeout))
	defer cancel()
	exists, err := as.authorRepository.CheckUsernameExists(c, username)
	if err != nil {
        return false, err
    }
    return exists, nil
}

// Create implements domain.AuthorService.
func (as *authorService) Create(ctx context.Context, author *domain.Author) error {
	if !isAlpha(author.Username) {
		return domain.ErrInvalidUsername
	}

	author.CreatedAt = time.Now()
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()
	return as.authorRepository.Create(c, author)
}

// Fetch implements domain.AuthorService.
func (as *authorService) Fetch(ctx context.Context) ([]domain.AuthorResponse, error) {
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()

	authors, err := as.authorRepository.Fetch(c)
	if err != nil {
		return nil, err
	}

	return authors, nil

}

// GetByID implements domain.AuthorService.
func (as *authorService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Author, error) {
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()

	author, err := as.authorRepository.GetByID(c, id)
	if err != nil {
		return nil, domain.ErrAuthorNotFound
	}

	return author, nil
}

// GetByUsername implements domain.AuthorService.
func (as *authorService) GetByUsername(ctx context.Context, username string) (*domain.Author, error) {
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()

	author, err := as.authorRepository.GetByUsername(c, username)
	if err != nil {
		return nil, domain.ErrAuthorNotFound
	}

	return author, nil
}

func NewAuthorService(authorRepository domain.AuthorRepository, timeout time.Duration) domain.AuthorService {
	return &authorService{
		authorRepository: authorRepository,
		contextTimeout:   timeout,
	}
}
