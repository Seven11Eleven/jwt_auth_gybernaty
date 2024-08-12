package domain

import (
	"context"
	"time"
)

type Article struct {
	ID        string     `json:"id"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleRepository interface {
	Create(ctx context.Context, article *Article) error
	FetchByUserID(ctx context.Context, userID string) ([]Article, error)
	GetByID(ctx context.Context, artID string) (*Article, error)
}

type ArticleService interface{
	Create(ctx context.Context, article *Article) error
	FetchByUserID(ctx context.Context, userID string) ([]Article, error)
	GetByID(ctx context.Context, artID string) (*Article, error)
}
