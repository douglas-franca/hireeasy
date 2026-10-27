package bootstrap

import (
	"context"
	"os"
	"strconv"
	"time"

	apphttp "github.com/douglas-franca/hireeasy/internal/app/http"
	companiesapp "github.com/douglas-franca/hireeasy/internal/modules/companies/app"
	companiespostgres "github.com/douglas-franca/hireeasy/internal/modules/companies/infra/postgres"
	companieshttp "github.com/douglas-franca/hireeasy/internal/modules/companies/transport/http"
	sessionsapp "github.com/douglas-franca/hireeasy/internal/modules/sessions/app"
	sessionspostgres "github.com/douglas-franca/hireeasy/internal/modules/sessions/infra/postgres"
	sessionshttp "github.com/douglas-franca/hireeasy/internal/modules/sessions/transport/http"
	"github.com/douglas-franca/hireeasy/internal/platform/email"
	"github.com/douglas-franca/hireeasy/internal/platform/postgres"

	usersapp "github.com/douglas-franca/hireeasy/internal/modules/users/app"
	userspostgres "github.com/douglas-franca/hireeasy/internal/modules/users/infra/postgres"
	usershttp "github.com/douglas-franca/hireeasy/internal/modules/users/transport/http"
)

func NewApp(ctx context.Context) (*App, error) {
	db, err := postgres.NewPool(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	router := apphttp.NewRouter()

	api := router.Group("/api/v1")

	userRepo := userspostgres.NewRepository(db)
	verifyURL := os.Getenv("VERIFY_EMAIL_URL")
	mailer := email.NewResend(os.Getenv("RESEND_API_KEY"), os.Getenv("EMAIL_FROM"), verifyURL)
	issuer := usersapp.NewVerificationTokenIssuer(userRepo, mailer, 24*time.Hour)
	userHandler := usershttp.NewHandler(
		usersapp.NewRegisterUser(userRepo, os.Getenv("PASSWORD_PEPPER"), issuer),
		usersapp.NewGetCurrentUser(userRepo),
		usersapp.NewUpdateCurrentUser(userRepo),
		usersapp.NewDeleteCurrentUser(userRepo, userRepo),
		usersapp.NewRequestEmailVerification(userRepo, issuer),
		usersapp.NewConfirmEmailVerification(userRepo),
	)
	duration := 7 * 24 * time.Hour
	if hours, err := strconv.Atoi(os.Getenv("SESSION_DURATION_HOURS")); err == nil && hours > 0 {
		duration = time.Duration(hours) * time.Hour
	}
	secureCookie := os.Getenv("COOKIE_SECURE") == "true"
	sessionRepo := sessionspostgres.NewRepository(db)
	createSession := sessionsapp.NewCreateSession(userRepo, sessionRepo, os.Getenv("PASSWORD_PEPPER"), duration)
	deleteSession := sessionsapp.NewDeleteSession(sessionRepo)
	authenticateSession := sessionsapp.NewAuthenticateSession(sessionRepo)
	sessionHandler := sessionshttp.NewHandler(createSession, deleteSession, secureCookie, duration)
	auth := sessionshttp.Required(authenticateSession)
	companyRepo := companiespostgres.NewRepository(db)
	companyHandler := companieshttp.NewHandler(companiesapp.NewCreateCompany(companyRepo), companiesapp.NewListCompanies(companyRepo), companiesapp.NewGetCompany(companyRepo))

	usershttp.RegisterRoutes(api, userHandler, auth)
	api.POST("/sessions", sessionHandler.Create)
	api.DELETE("/sessions/current", sessionHandler.Delete)
	companieshttp.RegisterRoutes(api, companyHandler, auth)

	return &App{
		Router: router,
		close:  db.Close,
	}, nil
}
