package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	"github.com/google/uuid"
)

type GetCompany struct{ repository domain.Repository }

func NewGetCompany(repository domain.Repository) *GetCompany {
	return &GetCompany{repository: repository}
}
func (u *GetCompany) Execute(ctx context.Context, userID, companyID uuid.UUID) (*domain.Company, error) {
	return u.repository.FindForUser(ctx, companyID, userID)
}
