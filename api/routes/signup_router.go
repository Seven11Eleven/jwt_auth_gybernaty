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

func NewSignupRouter(env *config.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	aur := repository.NewAuthorPgxRepository(db)
	sc := controller.SignUpController{
		SignUpService: service.NewSignupService(aur, timeout),
		Env : env,
	}
	group.POST("/signup", sc.SignUp)
}
