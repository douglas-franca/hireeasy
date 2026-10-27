package app

import (
	"context"
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	"github.com/google/uuid"
	"testing"
)

type companyRepository struct {
	owner uuid.UUID
	name  string
}

func (r *companyRepository) CreateForVerifiedUser(_ context.Context, owner uuid.UUID, name string) (*domain.Company, error) {
	r.owner = owner
	r.name = name
	return &domain.Company{ID: uuid.New(), Name: name, Role: "owner"}, nil
}
func (r *companyRepository) ListByUser(context.Context, uuid.UUID) ([]domain.Company, error) {
	return nil, nil
}
func (r *companyRepository) FindForUser(context.Context, uuid.UUID, uuid.UUID) (*domain.Company, error) {
	return nil, domain.ErrCompanyNotFound
}
func TestCreateCompanyUsesAuthenticatedOwner(t *testing.T) {
	repository := &companyRepository{}
	owner := uuid.New()
	company, err := NewCreateCompany(repository).Execute(context.Background(), owner, CreateCompanyInput{Name: "Acme"})
	if err != nil {
		t.Fatal(err)
	}
	if repository.owner != owner || repository.name != "Acme" || company.Role != "owner" {
		t.Fatal("expected owner membership")
	}
}
