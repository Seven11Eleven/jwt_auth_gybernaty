package routes

import (
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/api/middleware"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(env config.Env, timeout time.Duration, db *pgx.Conn, gin *gin.Engine){
	publicRouter := gin.Group("")

	NewSignupRouter(&env, timeout, db, publicRouter)
	NewLoginRouter(&env, timeout, db, publicRouter)
	NewRefreshTokenRouter(&env, timeout, db, *publicRouter)

	protectedRouter := gin.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTSecret))

	NewAuthorRouter(&env, timeout, db, protectedRouter)
	NewArticleRouter(&env, timeout, db, protectedRouter)

	
}