package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	"github.com/google/uuid"
)

type CreateCompanyInput struct{ Name string }
type CreateCompany struct{ repository domain.Repository }

func NewCreateCompany(repository domain.Repository) *CreateCompany {
	return &CreateCompany{repository: repository}
}
func (u *CreateCompany) Execute(ctx context.Context, userID uuid.UUID, input CreateCompanyInput) (*domain.Company, error) {
	return u.repository.CreateForVerifiedUser(ctx, userID, input.Name)
}
