package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"name"`
	Articles  []Article `json:"articles"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt time.Time `json:"updated_at"`
}

type AuthorResponse struct {
	Username string            `json:"username"`
	Articles []ArticleResponse `json:"articles"`
}

type AuthorRepository interface {
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	Create(ctx context.Context, author *Author) error
	Fetch(ctx context.Context) ([]AuthorResponse, error)
	GetByUsername(ctx context.Context, username string) (*Author, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Author, error)
}

type AuthorService interface {
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	Create(ctx context.Context, author *Author) error
	Fetch(ctx context.Context) ([]AuthorResponse, error)
	GetByUsername(ctx context.Context, username string) (*Author, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Author, error)
}
