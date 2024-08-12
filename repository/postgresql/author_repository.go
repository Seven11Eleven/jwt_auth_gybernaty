package repository

import (
	"context"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/jackc/pgx/v5"
)

type authorPgxRepository struct {
	db *pgx.Conn
}

func NewAuthorPgxRepository(db *pgx.Conn) domain.AuthorRepository {
	return &authorPgxRepository{db: db}
}

// Create implements domain.AuthorRepository.
func (ar *authorPgxRepository) Create(ctx context.Context, author *domain.Author) error {
	query := `INSERT INTO authors (id, username, password, created_at) VALUES($1, $2, $3, $4, $5)`
	_, err := ar.db.Exec(ctx, query, author.ID, author.Username, author.Password, author.CreatedAt)
	return err
}

// Fetch implements domain.AuthorRepository.
func (ar *authorPgxRepository) Fetch(ctx context.Context) ([]domain.Author, error) {
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

	authorsMap := make(map[string]*domain.Author)
	for rows.Next() {
		for rows.Next() {
			var authorID, articleID, authorUsername, articleTitle, articleContent string

			err := rows.Scan(&authorID, &authorUsername, &articleID, &articleTitle, &articleContent)
			if err != nil {
				return nil, err
			}
			if author, exists := authorsMap[authorID]; exists {
				author.Articles = append(author.Articles, domain.Article{
					ID:      articleID,
					Title:   articleTitle,
					Content: articleContent,
				})
			} else {
				authorsMap[authorID] = &domain.Author{
					ID:       authorID,
					Username: authorUsername,
					Articles: []domain.Article{
						{
							ID:      articleID,
							Title:   articleTitle,
							Content: articleContent,
						},
					},
				}
			}
		}
	}

	var authors []domain.Author
	for _, author := range authorsMap {
		authors = append(authors, *author)
	}

	return authors, nil
}

// GetByUsername implements domain.AuthorRepository.
func (ar *authorPgxRepository) GetByUsername(ctx context.Context, username string) (*domain.Author, error) {
	query := `
		SELECT
			a.id, a.username, ar.id, ar.title, ar.content
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
		var authorID, authorUsername, articleID, articleTitle, articleContent string

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
		
		if articleID != ""{
			author.Articles = append(author.Articles, domain.Article{
				ID: articleID,
				Title: articleTitle,
				Content: articleContent,
			})
		}
	}

	if author == nil{
		return nil, domain.ErrAuthorNotFound
	}

	return author, nil
}
