package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type DeleteCurrentUser struct {
	ownership domain.OwnershipChecker
	deleter   domain.Deleter
}

func NewDeleteCurrentUser(ownership domain.OwnershipChecker, deleter domain.Deleter) *DeleteCurrentUser {
	return &DeleteCurrentUser{ownership: ownership, deleter: deleter}
}

func (u *DeleteCurrentUser) Execute(ctx context.Context, userID uuid.UUID) error {
	owns, err := u.ownership.HasOwnedCompanies(ctx, userID)
	if err != nil {
		return err
	}

	if owns {
		return domain.ErrUserOwnsCompany
	}

	return u.deleter.Delete(ctx, userID)
}
