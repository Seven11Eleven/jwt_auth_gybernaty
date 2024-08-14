package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain/mocks"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchByUserID(t *testing.T) {
	mockArticleRepository := new(mocks.ArticleRepository)
	authorUUID := uuid.New()

	t.Run("success", func(t *testing.T) {

		mockArticle := domain.ArticleResponse{
			Title:   "Test Title",
			Content: "i am rockstar",
		}

		mockListArticle := make([]domain.ArticleResponse, 0)
		mockListArticle = append(mockListArticle, mockArticle)

		mockArticleRepository.On("FetchByUserID", mock.Anything, authorUUID).Return(mockListArticle, nil).Once()

		serv := service.NewArticleService(mockArticleRepository, time.Second*2)

		list, err := serv.FetchByUserID(context.Background(), authorUUID)

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListArticle))

		mockArticleRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockArticleRepository.On("FetchByUserID", mock.Anything, authorUUID).Return(nil, errors.New("Unexpected")).Once()

		serv := service.NewArticleService(mockArticleRepository, time.Second*2)

		list, err := serv.FetchByUserID(context.Background(), authorUUID)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockArticleRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockArticleRepository := new(mocks.ArticleRepository)
	serv := service.NewArticleService(mockArticleRepository, time.Second*2)

	t.Run("success", func(t *testing.T) {
		mockArticle := &domain.Article{
			Title:   "Valid Title",
			Content: "Valid content",
		}

		mockArticleRepository.On("Create", mock.Anything, mockArticle).Return(nil).Once()

		err := serv.Create(context.Background(), mockArticle)

		assert.NoError(t, err)
		mockArticleRepository.AssertExpectations(t)
	})

	t.Run("invalid title lenght", func(t *testing.T) {
		mockArticle := &domain.Article{
			Title:   "ltf", //ну типа меньше четырех ltf - less than four =)
			Content: "Valid content",
		}

		err := serv.Create(context.Background(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidTitle, err)
		mockArticleRepository.AssertNotCalled(t, "Create", mock.Anything, mockArticle)
	})

	t.Run("invalid title characters", func(t *testing.T) {
		mockArticle := &domain.Article{
			Title:   "Невалидный1337#Заголовок",
			Content: "Valid content",
		}

		err := serv.Create(context.Background(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidTitle, err)
		mockArticleRepository.AssertNotCalled(t, "Create", mock.Anything, mockArticle)
	})

	t.Run("invalid content characters", func(t *testing.T) {
		mockArticle := &domain.Article{
			Title:   "Valid Title",
			Content: "невалидное содержимое 1337 @#$#$%,.]|",
		}

		err := serv.Create(context.Background(), mockArticle)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidContent, err)
		mockArticleRepository.AssertNotCalled(t, "Create", mock.Anything, mockArticle)
	})
}

func TestGetByID(t *testing.T) {
	mockArticleRepository := new(mocks.ArticleRepository)
	serv := service.NewArticleService(mockArticleRepository, time.Second*2)
	articleUUID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockArticle := &domain.ArticleResponse{
			Title:   "Test Title",
			Content: "Test content",
		}

		mockArticleRepository.On("GetByID", mock.Anything, articleUUID).Return(mockArticle, nil).Once()

		article, err := serv.GetByID(context.Background(), articleUUID)

		assert.NoError(t, err)
		assert.NotNil(t, article)
		assert.Equal(t, mockArticle, article)
		mockArticleRepository.AssertExpectations(t)
	})

	t.Run("article not found", func(t *testing.T) {
		mockArticleRepository.On("GetByID", mock.Anything, articleUUID).Return(nil, domain.ErrArticleNotFound).Once()

		article, err := serv.GetByID(context.Background(), articleUUID)

		assert.Error(t, err)
		assert.Nil(t, article)
		assert.Equal(t, domain.ErrArticleNotFound, err)
		mockArticleRepository.AssertExpectations(t)
	})
}
