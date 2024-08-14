package routes

import (
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/api/controller"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"

	repository "github.com/Seven11Eleven/jwt_auth_gybernaty/repository/postgresql"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewArticleRouter(env *config.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	ar := repository.NewArticlePgxRepository(db)
	arc := &controller.ArticleController{
		ArticleService: service.NewArticleService(ar, timeout),
		Env:            env,
	}
	group.GET("/article/:id", arc.GetByID)
	group.POST("/article", arc.Create)
}
