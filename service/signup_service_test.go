package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain/mocks"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUpService_CheckUsernameExists(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	authorUsername := "vasil"

	t.Run("exists", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(true, nil).Once()

		serv := service.NewSignupService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.NoError(t, err)
		assert.True(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("not exists", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(false, nil).Once()

		serv := service.NewSignupService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.NoError(t, err)
		assert.False(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockAuthorRepository.On("CheckUsernameExists", mock.Anything, authorUsername).Return(false, errors.New("Unexpected error")).Once()

		serv := service.NewSignupService(mockAuthorRepository, time.Second*2)

		exists, err := serv.CheckUsernameExists(context.Background(), authorUsername)

		assert.Error(t, err)
		assert.False(t, exists)

		mockAuthorRepository.AssertExpectations(t)
	})
}

func TestSignUpService_Create(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)

	t.Run("success", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: "vasya",
		}

		mockAuthorRepository.On("Create", mock.Anything, mockAuthor).Return(nil).Once()

		serv := service.NewSignupService(mockAuthorRepository, time.Second*2)

		err := serv.Create(context.Background(), mockAuthor)

		assert.NoError(t, err)
		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("invalid username", func(t *testing.T) {
		mockAuthor := &domain.Author{
			Username: "Василий В. В.",
		}

		serv := service.NewSignupService(mockAuthorRepository, time.Second*2)

		err := serv.Create(context.Background(), mockAuthor)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrInvalidUsername, err)
	})
}
