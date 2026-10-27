package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/douglas-franca/hireeasy/internal/modules/sessions/domain"
	sessionssqlc "github.com/douglas-franca/hireeasy/internal/modules/sessions/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct{ queries *sessionssqlc.Queries }

func NewRepository(db sessionssqlc.DBTX) *Repository {
	return &Repository{queries: sessionssqlc.New(db)}
}
func (r *Repository) Create(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error {
	return r.queries.CreateSession(ctx, sessionssqlc.CreateSessionParams{
		UserID:    pgtype.UUID{Bytes: userID, Valid: true},
		TokenHash: tokenHash,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})
}
func (r *Repository) FindUserID(ctx context.Context, tokenHash string) (uuid.UUID, error) {
	value, err := r.queries.FindSessionUserID(ctx, tokenHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return uuid.Nil, domain.ErrUnauthenticated
	}
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.UUID(value.Bytes), nil
}
func (r *Repository) Delete(ctx context.Context, tokenHash string) error {
	return r.queries.DeleteSession(ctx, tokenHash)
}
