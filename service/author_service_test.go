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

func TestCheckUsernameExists(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	authorUsername := "vasyapupkin"

	t.Run("exists", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(true, nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.NoError(t, err)
		assert.True(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("not exists", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(false, nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.NoError(t, err)
		assert.False(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(false, errors.New("Unexpected error")).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.Error(t, err)
		assert.False(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})
}

func TestCreateAuthor(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)

	t.Run("success", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: "vasyapupkin",
		}

		mockAuthorRepository.On("Create", mock.Anything, mockAuthor).Return(nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		err := serv.Create(context.Background(), mockAuthor)

		assert.NoError(t, err)
		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("invalid username", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: "вася321421#$@$",
		}

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		err := serv.Create(context.Background(), mockAuthor)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidUsername, err)
	})
}

func TestFetchAuthors(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)

	t.Run("success", func(t *testing.T) {
		mockAuthor := domain.AuthorResponse{
			Username: "vasyapupkin",
		}

		mockListAuthor := make([]domain.AuthorResponse, 0)
		mockListAuthor = append(mockListAuthor, mockAuthor)

		mockAuthorRepository.On("Fetch", mock.Anything).Return(mockListAuthor, nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		authors, err := serv.Fetch(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, authors)
		assert.Len(t, authors, len(mockListAuthor))

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockAuthorRepository.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		authors, err := serv.Fetch(context.Background())

		assert.Error(t, err)
		assert.Nil(t, authors)

		mockAuthorRepository.AssertExpectations(t)
	})
}

func TestGetAuthorByID(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	authorID := "b951ad2b-d011-4905-8a8b-d3a90889d876"
	authorUUID := uuid.MustParse(authorID)

	t.Run("success", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: "vasiliy",
		}

		mockAuthorRepository.On("GetByID", mock.Anything, authorUUID).Return(mockAuthor, nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		author, err := serv.GetByID(context.Background(), authorUUID)

		assert.NoError(t, err)
		assert.NotNil(t, author)
		assert.Equal(t, mockAuthor.Username, author.Username)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("wrong id provided", func(t *testing.T) {
		mockAuthorRepository.On("GetByID", mock.Anything, authorUUID).Return(nil, domain.ErrAuthorNotFound).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		author, err := serv.GetByID(context.Background(), authorUUID)

		assert.Error(t, err)
		assert.Nil(t, author)
		assert.Equal(t, domain.ErrAuthorNotFound, err)

		mockAuthorRepository.AssertExpectations(t)
	})
}

func TestGetAuthorByUsername(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	username := "vasiliy"

	t.Run("success", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: username,
		}

		mockAuthorRepository.On("GetByID", mock.Anything, username).Return(mockAuthor, nil).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		author, err := serv.GetByUsername(context.Background(), username)

		assert.NoError(t, err)
		assert.NotNil(t, author)
		assert.Equal(t, mockAuthor.Username, author.Username)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("wrong username provid ed", func(t *testing.T) {
		mockAuthorRepository.On("GetByUsername", mock.Anything, username).Return(nil, domain.ErrAuthorNotFound).Once()

		serv := service.NewAuthorService(mockAuthorRepository, time.Second*2)

		author, err := serv.GetByUsername(context.Background(), username)

		assert.Error(t, err)
		assert.Nil(t, author)
		assert.Equal(t, domain.ErrAuthorNotFound, err)

		mockAuthorRepository.AssertExpectations(t)
	})
}
