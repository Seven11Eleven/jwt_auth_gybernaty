package app

import (
	"context"
	"fmt"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/api/routes"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/config"
	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type App struct {
	Server *gin.Engine
	DB     *pgx.Conn
	Env    *config.Env
}

func NewApp(ctx context.Context) (*App, error) {
	env := config.NewEnv()

	db := database.NewPostgreSQLConnection(env)

	ginEngine := gin.Default()

	
	routes.SetupRoutes(*env, 10*time.Second, db, ginEngine)

	return &App{
		Server: ginEngine,
		DB:     db,
		Env:    env,
	}, nil

		
}

func(a *App) Run() error{
	port := a.Env.ServerPort
	if port == ""{
		fmt.Println("походу что то не загрузилось")
		port = "8080"
	}

	return a.Server.Run(fmt.Sprintf(":%s", port))
}


func (a *App) Close(){
	database.ClosePostgreSQLConnection(a.DB)
}