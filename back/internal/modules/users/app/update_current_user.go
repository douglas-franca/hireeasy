package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type UpdateCurrentUserInput struct{ Name string }
type UpdateCurrentUser struct{ updater domain.NameUpdater }

func NewUpdateCurrentUser(updater domain.NameUpdater) *UpdateCurrentUser {
	return &UpdateCurrentUser{updater: updater}
}

func (u *UpdateCurrentUser) Execute(ctx context.Context, userID uuid.UUID, input UpdateCurrentUserInput) error {
	return u.updater.UpdateName(ctx, userID, input.Name)
}
