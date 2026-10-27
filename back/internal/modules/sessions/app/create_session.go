package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/sessions/domain"
	users "github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/douglas-franca/hireeasy/internal/platform/security/password"
	"time"
)

type CreateSessionInput struct {
	Email    string
	Password string
}
type CreateSession struct {
	users    users.ByEmailFinder
	sessions domain.Repository
	pepper   string
	duration time.Duration
}

func NewCreateSession(users users.ByEmailFinder, sessions domain.Repository, pepper string, duration time.Duration) *CreateSession {
	return &CreateSession{users: users, sessions: sessions, pepper: pepper, duration: duration}
}
func (u *CreateSession) Execute(ctx context.Context, input CreateSessionInput) (string, error) {
	user, err := u.users.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}
	if password.Compare(user.PasswordHash, input.Password, u.pepper) != nil {
		return "", domain.ErrInvalidCredentials
	}
	token, err := newToken()
	if err != nil {
		return "", err
	}
	if err := u.sessions.Create(ctx, user.ID, hashToken(token), time.Now().Add(u.duration)); err != nil {
		return "", err
	}
	return token, nil
}
