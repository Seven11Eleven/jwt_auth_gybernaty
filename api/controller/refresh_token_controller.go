package controller

import (
	"net/http"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenService domain.RefreshTokenService
	Env                 *config.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var req domain.RefreshTokenRequest

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := rtc.RefreshTokenService.ExtractIDFromToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	author, err := rtc.RefreshTokenService.GetAuthorByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := rtc.RefreshTokenService.CreateAccessToken(author, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := rtc.RefreshTokenService.CreateRefreshToken(author, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
