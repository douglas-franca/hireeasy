package postgres

import (
	"context"
	"errors"

	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	companiessqlc "github.com/douglas-franca/hireeasy/internal/modules/companies/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct{ queries *companiessqlc.Queries }

func NewRepository(db companiessqlc.DBTX) *Repository {
	return &Repository{queries: companiessqlc.New(db)}
}

func (r *Repository) CreateForVerifiedUser(ctx context.Context, userID uuid.UUID, name string) (*domain.Company, error) {
	row, err := r.queries.CreateCompanyForVerifiedUser(ctx, companiessqlc.CreateCompanyForVerifiedUserParams{Name: name, UserID: uuidValue(userID)})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrEmailNotVerified
	}
	if err != nil {
		return nil, err
	}
	return companyFromCreateRow(row), nil
}

func (r *Repository) ListByUser(ctx context.Context, userID uuid.UUID) ([]domain.Company, error) {
	rows, err := r.queries.ListCompaniesByUser(ctx, uuidValue(userID))
	if err != nil {
		return nil, err
	}
	companies := make([]domain.Company, 0, len(rows))
	for _, row := range rows {
		companies = append(companies, companyFromListRow(row))
	}
	return companies, nil
}

func (r *Repository) FindForUser(ctx context.Context, companyID, userID uuid.UUID) (*domain.Company, error) {
	row, err := r.queries.FindCompanyForUser(ctx, companiessqlc.FindCompanyForUserParams{ID: uuidValue(companyID), UserID: uuidValue(userID)})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCompanyNotFound
	}
	if err != nil {
		return nil, err
	}
	company := companyFromFindRow(row)
	return &company, nil
}

func uuidValue(id uuid.UUID) pgtype.UUID { return pgtype.UUID{Bytes: id, Valid: true} }
