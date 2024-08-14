package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID     `json:"id"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
} 

type ArticleResponse struct{
	Title string `json:"title"`
	Content string `json:"content"`
}

type ArticleRepository interface {
	Create(ctx context.Context, article *Article) error
	FetchByUserID(ctx context.Context, userID uuid.UUID) ([]ArticleResponse, error)
	GetByID(ctx context.Context, artID uuid.UUID) (*ArticleResponse, error)
}

type ArticleService interface{
	Create(ctx context.Context, article *Article) error
	FetchByUserID(ctx context.Context, userID uuid.UUID) ([]ArticleResponse, error)
	GetByID(ctx context.Context, artID uuid.UUID) (*ArticleResponse, error)
}
