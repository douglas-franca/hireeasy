package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	"github.com/google/uuid"
)

type ListCompanies struct{ repository domain.Repository }

func NewListCompanies(repository domain.Repository) *ListCompanies {
	return &ListCompanies{repository: repository}
}
func (u *ListCompanies) Execute(ctx context.Context, userID uuid.UUID) ([]domain.Company, error) {
	return u.repository.ListByUser(ctx, userID)
}
