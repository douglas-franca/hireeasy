package postgres

import (
	"context"

	"github.com/douglas-franca/hionshop/internal/modules/users/domain"
	usersqlc "github.com/douglas-franca/hionshop/internal/modules/users/infra/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *usersqlc.Queries
}

func NewRepository(db usersqlc.DBTX) *Repository {
	return &Repository{
		q: usersqlc.New(db),
	}
}

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	return r.q.CreateUser(ctx, usersqlc.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	})
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.q.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return toDomainUser(&user), nil
}
