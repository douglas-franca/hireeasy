package app

import (
	"context"
	"testing"
	"time"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type registrationRepository struct {
	user      *domain.User
	tokenHash string
}

func (r *registrationRepository) Create(_ context.Context, user *domain.User) error {
	r.user = user
	return nil
}

func (r *registrationRepository) FindByID(context.Context, uuid.UUID) (*domain.User, error) {
	return r.user, nil
}

func (r *registrationRepository) SetVerificationToken(_ context.Context, _ uuid.UUID, hash string, _ time.Time) error {
	r.tokenHash = hash
	return nil
}

type emailSender struct {
	to    string
	token string
}

func (s *emailSender) SendVerification(_ context.Context, to, token string) error {
	s.to = to
	s.token = token
	return nil
}

func TestRegisterUserHashesPasswordAndSendsVerification(t *testing.T) {
	repository := &registrationRepository{}
	sender := &emailSender{}
	issuer := NewVerificationTokenIssuer(repository, sender, time.Hour)
	useCase := NewRegisterUser(repository, "pepper", issuer)
	if err := useCase.Execute(context.Background(), RegisterUserInput{Name: "Ana", Email: "ana@example.com", Password: "password123"}); err != nil {
		t.Fatal(err)
	}
	if repository.user.PasswordHash == "password123" {
		t.Fatal("password must be hashed")
	}
	if sender.to != "ana@example.com" || sender.token == "" || repository.tokenHash == "" {
		t.Fatal("verification token must be stored and sent")
	}
}
