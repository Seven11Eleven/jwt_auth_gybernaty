package service_test

import (
	"context"

	"testing"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain/mocks"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByUsername(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	mockAuthor := &domain.Author{
		Username: "Vasiliy",
	}

	t.Run("success", func(t *testing.T) {
		mockAuthorRepository.On("GetByUsername", mock.Anything, mockAuthor.Username).Return(mockAuthor, nil).Once()

		serv := service.NewLoginService(mockAuthorRepository, time.Second*2)
		author, err := serv.GetUserByUsername(context.Background(), mockAuthor.Username)

		assert.NoError(t, err)
		assert.NotNil(t, author)
		assert.Equal(t, mockAuthor.Username, author.Username)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("not exists", func(t *testing.T) {
		mockAuthorRepository.On("GetByUsername", mock.Anything, mockAuthor.Username).Return(nil, domain.ErrAuthorNotFound).Once()

		serv := service.NewLoginService(mockAuthorRepository, time.Second*2)
		author, err := serv.GetUserByUsername(context.Background(), mockAuthor.Username)

		assert.Error(t, err)
		assert.Nil(t, author)

		mockAuthorRepository.AssertExpectations(t)
	})
}
