package service

import (
	"context"
	"regexp"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
)

func isAlpha(str string) bool {
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, str)
	return match
}

type articleService struct {
	articleRepository domain.ArticleRepository
	contextTimeout    time.Duration
}

func NewArticleService(articleRepository domain.ArticleRepository, timeout time.Duration) domain.ArticleService {
	return &articleService{
		articleRepository: articleRepository,
		contextTimeout:    timeout,
	}
}

// Create implements domain.ArticleService.
func (as *articleService) Create(ctx context.Context, article *domain.Article) error {
	if len(article.Title) < 4 || len(article.Title) > 100 {
		return domain.ErrInvalidTitle
	}

	if !isAlpha(article.Title) {
		return domain.ErrInvalidTitle
	}

	if !isAlpha(article.Content) {
		return domain.ErrInvalidContent
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()
	return as.articleRepository.Create(c, article)
}

// FetchByUserID implements domain.ArticleService.
func (as *articleService) FetchByUserID(ctx context.Context, userID string) ([]domain.Article, error) {
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()

	articles, err := as.articleRepository.FetchByUserID(c, userID)
	if err != nil {
		return nil, domain.ErrAuthorNotFound
	}

	if len(articles) == 0 {
		return nil, domain.ErrArticleNotFound
	}

	return articles, nil
}

// GetByID implements domain.ArticleService.
func (as *articleService) GetByID(ctx context.Context, artID string) (*domain.Article, error) {
	c, cancel := context.WithTimeout(ctx, as.contextTimeout)
	defer cancel()

	article, err := as.articleRepository.GetByID(c, artID)
	if err != nil {
		return nil, domain.ErrArticleNotFound
	}
	return article, nil
}
