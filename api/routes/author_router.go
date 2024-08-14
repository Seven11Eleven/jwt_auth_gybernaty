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

func NewAuthorRouter(env *config.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	aur := repository.NewAuthorPgxRepository(db)
	auc := &controller.AuthorController{
		AuthorService: service.NewAuthorService(aur, timeout),
		Env:           env,
	}
	group.GET("/all", auc.Fetch)
}
