package routes

import (
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(env *app.Env, timeout time.Duration, db *pgx.Conn, gin *gin.Engine){
	publicRouter := gin.Group("")

	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter()

	protectedRouter := gin.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTSecret))

	NewAuthorRouter(env, timeout, db, protectedRouter)
	NewArticleRouter(env, timeout, db, protectedRouter)
}