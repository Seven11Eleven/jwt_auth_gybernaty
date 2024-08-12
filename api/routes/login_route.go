package routes

import (
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/app"
	repository "github.com/Seven11Eleven/jwt_auth_gybernaty/repository/postgresql"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewLoginRouter(env *app.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	aur := repository.NewAuthorPgxRepository(db)
	lc := &controller.LoginController{
		LoginService: service.NewLoginService(aur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}