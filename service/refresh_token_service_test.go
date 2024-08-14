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

func TestRefreshTokenService_GetAuthorByID(t *testing.T) {
	mockAuthorRepository := new(mocks.AuthorRepository)
	mockAuthor := &domain.Author{Username: "vasil"}
	authorID := "226a1439-85d0-401d-a942-eb17793465bd"
	authorUUID := uuid.MustParse(authorID)

	t.Run("success", func(t *testing.T) {
		mockAuthorRepository.On("GetByID", mock.Anything, authorUUID).Return(mockAuthor, nil).Once()

		serv := service.NewRefreshTokenService(mockAuthorRepository, time.Second*2)
		author, err := serv.GetAuthorByID(context.Background(), authorUUID)

		assert.NoError(t, err)
		assert.NotNil(t, author)
		assert.Equal(t, mockAuthor.Username, author.Username)

		mockAuthorRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockAuthorRepository.On("GetByID", mock.Anything, authorUUID).Return(nil, errors.New("Unexpected error")).Once()

		serv := service.NewRefreshTokenService(mockAuthorRepository, time.Second*2)
		author, err := serv.GetAuthorByID(context.Background(), authorUUID)

		assert.Error(t, err)
		assert.Nil(t, author)

		mockAuthorRepository.AssertExpectations(t)
	})
}

func TestRefreshTokenService_ExtractIDFromToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjIyNmExNDM5LTg1ZDAtNDAxZC1hOTQyLWViMTc3OTM0NjViZCIsImV4cCI6MTcyNDI4OTA1N30.frzUFSo0k0diX9VWxMlKIGh2lO2GCPcdyrRU7mE3DoI"
		expectedID := "226a1439-85d0-401d-a942-eb17793465bd"

		utilsMock := new(mocks.JWTUtils)
		utilsMock.On("ExtractIDFromToken", token).Return(expectedID, nil).Once()

		rts := service.NewRefreshTokenService(nil, time.Second*2)
		id, err := rts.ExtractIDFromToken(token)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
	})

	t.Run("invalid token", func(t *testing.T) {
		token := "invalid_token"
		utilsMock := new(mocks.JWTUtils)
		utilsMock.On("ExtractIDFromToken", token).Return("", domain.ErrInvalidToken).Once()

		rts := service.NewRefreshTokenService(nil, time.Second*2)
		id, err := rts.ExtractIDFromToken(token)

		assert.Error(t, err)
		assert.Empty(t, id)
	})
}
