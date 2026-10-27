package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/sessions/domain"
)

type DeleteSession struct{ sessions domain.Repository }

func NewDeleteSession(sessions domain.Repository) *DeleteSession {
	return &DeleteSession{sessions: sessions}
}
func (u *DeleteSession) Execute(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}
	return u.sessions.Delete(ctx, hashToken(token))
}
