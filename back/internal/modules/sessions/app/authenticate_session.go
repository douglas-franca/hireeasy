package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/sessions/domain"
	"github.com/google/uuid"
)

type AuthenticateSession struct{ sessions domain.Repository }

func NewAuthenticateSession(sessions domain.Repository) *AuthenticateSession {
	return &AuthenticateSession{sessions: sessions}
}
func (u *AuthenticateSession) Execute(ctx context.Context, token string) (uuid.UUID, error) {
	if token == "" {
		return uuid.Nil, domain.ErrUnauthenticated
	}
	return u.sessions.FindUserID(ctx, hashToken(token))
}
