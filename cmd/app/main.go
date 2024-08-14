package main

import (
	"context"
	"log"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/app"
)

func main() {
	app, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("Failed to init app : %v", err)
	}

	defer app.Close()

	if err := app.Run(); err != nil{
		log.Fatalf("Failed to run app: %v", err)
	}
}
