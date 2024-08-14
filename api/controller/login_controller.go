package controller

import (
	// "log"
	"net/http"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService domain.LoginService
	Env          *config.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var req domain.LoginRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, err := lc.LoginService.GetUserByUsername(c, req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	
	
	localParam := lc.Env.LocalParam

	
	if err := utils.CompareHashAndPassword(author.Password, req.Password, author.Salt, localParam); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	

	accessToken, err := lc.LoginService.CreateAccessToken(author, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := lc.LoginService.CreateRefreshToken(author, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
