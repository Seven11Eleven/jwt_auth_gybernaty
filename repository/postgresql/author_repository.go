package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type authorPgxRepository struct {
	db *pgx.Conn
}

// GetByID implements domain.AuthorRepository.
func (ar *authorPgxRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Author, error) {
	query := `
		SELECT
			a.id, a.username,
			ar.id, ar.title, ar.content
		FROM
			authors a
		LEFT JOIN
			articles ar ON a.id = ar.author_id
		WHERE 
			a.id = $1
		`

	rows, err := ar.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var author *domain.Author

	for rows.Next() {
		var authorID, articleID uuid.UUID
		var authorUsername, articleTitle, articleContent string

		err := rows.Scan(&authorID, &authorUsername, &articleID, &articleTitle, &articleContent)
		if err != nil {
			return nil, err
		}
		if author == nil {
			author = &domain.Author{
				ID:       authorID,
				Username: authorUsername,
				Articles: []domain.Article{},
			}
		}

		if articleID != uuid.Nil {
			author.Articles = append(author.Articles, domain.Article{
				ID:      articleID,
				Title:   articleTitle,
				Content: articleContent,
			})
		}

	}
	if author == nil {
		return nil, domain.ErrAuthorNotFound
	}

	return author, nil
}

func NewAuthorPgxRepository(db *pgx.Conn) domain.AuthorRepository {
	return &authorPgxRepository{db: db}
}

// Create implements domain.AuthorRepository.
func (ar *authorPgxRepository) Create(ctx context.Context, author *domain.Author) error {
	query := `INSERT INTO authors (id, username, password, salt, created_at) VALUES($1, $2, $3, $4, $5)`
	author.CreatedAt = time.Now()
	_, err := ar.db.Exec(ctx, query, author.ID, author.Username, author.Password, author.Salt, author.CreatedAt)
	return err
}

// Fetch implements domain.AuthorRepository.
func (ar *authorPgxRepository) Fetch(ctx context.Context) ([]domain.AuthorResponse, error) {
	query := `
		SELECT 
			a.id, a.username,
			ar.id, ar.title, ar.content
		FROM
			authors a
		LEFT JOIN 
			articles ar ON a.id = ar.author_id
	`
	rows, err := ar.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authorsMap := make(map[string]*domain.AuthorResponse)
	for rows.Next() {
		var authorID, articleID uuid.UUID
		var authorUsername string
		var articleTitle, articleContent sql.NullString

		err := rows.Scan(&authorID, &authorUsername, &articleID, &articleTitle, &articleContent)
		if err != nil {
			return nil, err
		}

		authorIDStr := authorID.String()

		if author, exists := authorsMap[authorIDStr]; exists {
			if articleID != uuid.Nil {
				author.Articles = append(author.Articles, domain.ArticleResponse{
					Title:   articleTitle.String,
					Content: articleContent.String,
				})
			}
		} else {
			authorsMap[authorIDStr] = &domain.AuthorResponse{
				
				Username: authorUsername,
				Articles: []domain.ArticleResponse{},
			}
			if articleID != uuid.Nil {
				authorsMap[authorIDStr].Articles = append(authorsMap[authorIDStr].Articles, domain.ArticleResponse{
					Title:   articleTitle.String,
					Content: articleContent.String,
				})
			}
		}
	}

	var authors []domain.AuthorResponse
	for _, author := range authorsMap {
		authors = append(authors, *author)
	}

	return authors, nil
}

// GetByUsername implements domain.AuthorRepository.
func (ar *authorPgxRepository) GetByUsername(ctx context.Context, username string) (*domain.Author, error) {
	query := `
		SELECT
			a.id, a.password, a.username, a.salt, ar.id,  ar.title, ar.content
		FROM 
			authors a
		LEFT JOIN
			articles ar ON a.id = ar.author_id
		WHERE 
			a.username = $1
	`

	rows, err := ar.db.Query(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var author *domain.Author

	for rows.Next() {
		var authorID, articleID uuid.UUID
		var authorUsername, authorSalt, authorPassword string
		var articleTitle, articleContent sql.NullString

		err := rows.Scan(&authorID, &authorPassword, &authorUsername, &authorSalt, &articleID, &articleTitle, &articleContent)
		if err != nil {
			return nil, err
		}

		if author == nil {
			author = &domain.Author{
				ID:       authorID,
				Username: authorUsername,
				Articles: []domain.Article{},
				Password: authorPassword,
				Salt:     authorSalt,
			}
		}

		if articleID != uuid.Nil {
			author.Articles = append(author.Articles, domain.Article{
				ID:      articleID,
				Title:   articleTitle.String,
				Content: articleContent.String,
			})
		}
	}

	if author == nil {
		return nil, domain.ErrAuthorNotFound
	}

	return author, nil
}

func (ar *authorPgxRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM authors WHERE username = $1)`

	var exists bool
	err := ar.db.QueryRow(ctx, query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
