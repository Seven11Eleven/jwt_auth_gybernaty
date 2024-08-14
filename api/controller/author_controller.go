package controller

import (
	"net/http"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	AuthorService domain.AuthorService
	Env           *config.Env
}

func (auc *AuthorController) Fetch(c *gin.Context) {
	var authors []domain.AuthorResponse

	authors, err := auc.AuthorService.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}


