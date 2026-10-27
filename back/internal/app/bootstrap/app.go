package bootstrap

import (
	"context"
	"os"

	apphttp "github.com/douglas-franca/hionshop/internal/app/http"
	"github.com/douglas-franca/hionshop/internal/platform/postgres"

	usersapp "github.com/douglas-franca/hionshop/internal/modules/users/app"
	userspostgres "github.com/douglas-franca/hionshop/internal/modules/users/infra/postgres"
	usershttp "github.com/douglas-franca/hionshop/internal/modules/users/transport/http"
)

func NewApp(ctx context.Context) (*App, error) {
	db, err := postgres.NewPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	router := apphttp.NewRouter()

	api := router.Group("/api/v1")

	userRepo := userspostgres.NewRepository(db)
	userService := usersapp.NewService(userRepo, os.Getenv("PASSWORD_PEPPER"))
	userHandler := usershttp.NewHandler(userService)

	usershttp.RegisterRoutes(api, userHandler)

	return &App{
		Router: router,
	}, nil
}
