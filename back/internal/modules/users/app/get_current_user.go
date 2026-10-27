package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type GetCurrentUser struct{ finder domain.ByIDFinder }

func NewGetCurrentUser(finder domain.ByIDFinder) *GetCurrentUser {
	return &GetCurrentUser{finder: finder}
}

func (u *GetCurrentUser) Execute(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return u.finder.FindByID(ctx, userID)
}
