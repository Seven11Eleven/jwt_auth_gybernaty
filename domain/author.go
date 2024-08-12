package domain

import (
	"context"
	"time"
)

type Author struct {
	ID        string     `json:"id"`
	Username      string    `json:"name"`
	Articles  []Article `json:"articles"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type AuthorRepository interface {
	Create(ctx context.Context, user *Author) error
	Fetch(ctx context.Context) ([]Author, error)
	GetByUsername(ctx context.Context, username string) (*Author, error)
}
