package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	usersqlc "github.com/douglas-franca/hireeasy/internal/modules/users/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *usersqlc.Queries
}

func NewRepository(db usersqlc.DBTX) *Repository {
	return &Repository{
		queries: usersqlc.New(db),
	}
}

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	params := usersqlc.CreateUserParams{
		ID:           uuidValue(user.ID),
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}

	err := r.queries.CreateUser(ctx, params)
	if isEmailAlreadyExists(err) {
		return domain.ErrEmailAlreadyExists
	}

	return err
}

func (r *Repository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	user, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return toDomainUser(&user), nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.User, error) {
	user, err := r.queries.FindUserByID(ctx, uuidValue(id))
	if err != nil {
		return nil, err
	}

	return toDomainUser(&user), nil
}

func (r *Repository) UpdateName(
	ctx context.Context,
	id uuid.UUID,
	name string,
) error {
	params := usersqlc.UpdateUserNameParams{
		ID:   uuidValue(id),
		Name: name,
	}

	_, err := r.queries.UpdateUserName(ctx, params)

	return err
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.queries.DeleteUser(ctx, uuidValue(id))

	return err
}

func (r *Repository) HasOwnedCompanies(
	ctx context.Context,
	id uuid.UUID,
) (bool, error) {
	return r.queries.HasOwnedCompanies(ctx, uuidValue(id))
}

func (r *Repository) SetVerificationToken(
	ctx context.Context,
	id uuid.UUID,
	hash string,
	expiresAt time.Time,
) error {
	params := usersqlc.SetEmailVerificationTokenParams{
		ID: uuidValue(id),
		EmailVerificationTokenHash: pgtype.Text{
			String: hash,
			Valid:  true,
		},
		EmailVerificationExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	}

	_, err := r.queries.SetEmailVerificationToken(ctx, params)

	return err
}

func (r *Repository) VerifyEmail(
	ctx context.Context,
	tokenHash string,
) (*domain.User, error) {
	token := pgtype.Text{
		String: tokenHash,
		Valid:  true,
	}

	user, err := r.queries.VerifyEmail(ctx, token)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrVerificationTokenInvalid
	}
	if err != nil {
		return nil, err
	}

	return toDomainUser(&user), nil
}

func uuidValue(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}
