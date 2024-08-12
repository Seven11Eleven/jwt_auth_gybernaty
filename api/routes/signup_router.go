package routes

import (
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/app"
	repository "github.com/Seven11Eleven/jwt_auth_gybernaty/repository/postgresql"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewSignupRouter(env *app.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup){
	aur := repository.NewAuthorPgxRepository(db)
	sc := controller.SignupController{
		SignUpService: service.NewSignupService(aur, timeout),
	}
	group.POST("/signup", sc.Signup)
}