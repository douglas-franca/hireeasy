package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, uuid.UUID, string, time.Time) error
	FindUserID(context.Context, string) (uuid.UUID, error)
	Delete(context.Context, string) error
}
