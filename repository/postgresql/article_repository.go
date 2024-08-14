package repository

import (
	"context"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type articlePgxRepository struct {
	db *pgx.Conn
}

func NewArticlePgxRepository(db *pgx.Conn) domain.ArticleRepository {
	return &articlePgxRepository{
		db: db,
	}
}

// Create implements domain.ArticleRepository.
func (ar *articlePgxRepository) Create(ctx context.Context, article *domain.Article) error {
	query := `INSERT INTO articles (id, title, content, author_id, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := ar.db.Exec(ctx, query, article.ID, article.Title, article.Content, article.Author.ID, article.CreatedAt)
	return err
}

// FetchByUserID implements domain.ArticleRepository.
func (ar *articlePgxRepository) FetchByUserID(ctx context.Context, userID uuid.UUID) ([]domain.ArticleResponse, error) {
	query := `SELECT id, title, content, author_id, created_id FROM articles WHERE author_id = $1`

	rows, err := ar.db.Query(ctx, query, userID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var articles []domain.ArticleResponse
	for rows.Next(){
		var article domain.ArticleResponse
		err := rows.Scan(&article.Title, &article.Content)
		if err != nil{
			return nil, err
		}
		articles = append (articles, article)
	}
	return articles, nil

}

// GetByID implements domain.ArticleRepository.
func (ar *articlePgxRepository) GetByID(ctx context.Context, artID uuid.UUID) (*domain.ArticleResponse, error) {
	query := `SELECT title, content FROM articles WHERE id = $1`
	row := ar.db.QueryRow(ctx, query, artID)

	var article domain.ArticleResponse
	err := row.Scan(&article.Title , &article.Content)
	if err != nil{
		if err == pgx.ErrNoRows{
			return nil, domain.ErrArticleNotFound
		}
		return nil, err
	}

	return &article, nil
}
