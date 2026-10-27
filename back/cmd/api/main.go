package main

import (
	"context"
	"log"
	"os"

	"github.com/douglas-franca/hionshop/internal/app/bootstrap"
)

func main() {
	app, err := bootstrap.NewApp(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
