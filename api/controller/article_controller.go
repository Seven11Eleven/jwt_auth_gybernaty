package controller

import (
	// "log"
	"net/http"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArticleController struct {
	ArticleService domain.ArticleService
	Env            *config.Env
}

func (arc *ArticleController) GetByID(c *gin.Context) {
	artID := c.Param("id")
	
	var article *domain.ArticleResponse
	artUUID := uuid.MustParse(artID)

	article, err := arc.ArticleService.GetByID(c, artUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": domain.ErrArticleNotFound})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (arc *ArticleController) Create(c *gin.Context) {
	var article domain.Article

	err := c.ShouldBind(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorID := c.GetString("authorID")

	authorUUID, err := uuid.Parse(authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.ID = uuid.New()
	article.Author.ID = authorUUID
	err = arc.ArticleService.Create(c, &article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfullResponse{
		Message: "article created successfully!",
	})
}
