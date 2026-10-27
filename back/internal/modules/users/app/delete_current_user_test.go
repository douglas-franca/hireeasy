package app

import (
	"context"
	"testing"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type ownerRepository struct{ deleted bool }

func (r *ownerRepository) HasOwnedCompanies(context.Context, uuid.UUID) (bool, error) {
	return true, nil
}

func (r *ownerRepository) Delete(context.Context, uuid.UUID) error { r.deleted = true; return nil }

func TestDeleteCurrentUserReturnsOwnershipError(t *testing.T) {
	repository := &ownerRepository{}
	err := NewDeleteCurrentUser(repository, repository).Execute(context.Background(), uuid.New())
	if err != domain.ErrUserOwnsCompany {
		t.Fatalf("got %v", err)
	}
	if repository.deleted {
		t.Fatal("company owner must not be deleted")
	}
}
