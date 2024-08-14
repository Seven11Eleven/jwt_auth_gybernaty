package controller_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/api/controller"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain/mocks"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	mockSignUpService := new(mocks.SignUpService)

	env := &config.Env{LocalParam: "localParam"}
	signUpController := &controller.SignUpController{
		SignUpService: mockSignUpService,
		Env:           env,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := `{"username":"testuser","password":"testpass"}`
	c.Request = httptest.NewRequest("POST", "/signup", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(bytes.NewBufferString(reqBody))

	mockSignUpService.On("CheckUsernameExists", mock.Anything, "testuser").Return(false, nil)
	mockSignUpService.On("Create", mock.Anything, mock.AnythingOfType("*domain.Author")).Return(nil)

	signUpController.SignUp(c)

	
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "you signed up successfully!")


	mockSignUpService.AssertExpectations(t)
}

func TestSignUp_UsernameExists(t *testing.T) {
	mockSignUpService := new(mocks.SignUpService)

	
	env := &config.Env{LocalParam: "localParam"}
	signUpController := &controller.SignUpController{
		SignUpService: mockSignUpService,
		Env:           env,
	}

	 
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	
	reqBody := `{"username":"armansu","password":"testpass"}`
	c.Request = httptest.NewRequest("POST", "/signup", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(bytes.NewBufferString(reqBody))

	mockSignUpService.On("CheckUsernameExists", mock.Anything, "testuser").Return(true, nil)

	signUpController.SignUp(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "username already exists")


	mockSignUpService.AssertExpectations(t)
}

func TestSignUp_FailureOnCreate(t *testing.T) {
	mockSignUpService := new(mocks.SignUpService)

	env := &config.Env{LocalParam: "localParam"}
	signUpController := &controller.SignUpController{
		SignUpService: mockSignUpService,
		Env:           env,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := `{"username":"testuser","password":"testpass"}`
	c.Request = httptest.NewRequest("POST", "/signup", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(bytes.NewBufferString(reqBody))

	mockSignUpService.On("CheckUsernameExists", mock.Anything, "testuser").Return(false, nil)
	mockSignUpService.On("Create", mock.Anything, mock.AnythingOfType("*domain.Author")).Return(errors.New("failed to create user"))

	signUpController.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "failed to create user")

	mockSignUpService.AssertExpectations(t)
}
