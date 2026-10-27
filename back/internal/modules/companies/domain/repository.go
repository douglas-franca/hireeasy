package domain

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateForVerifiedUser(context.Context, uuid.UUID, string) (*Company, error)
	ListByUser(context.Context, uuid.UUID) ([]Company, error)
	FindForUser(context.Context, uuid.UUID, uuid.UUID) (*Company, error)
}
