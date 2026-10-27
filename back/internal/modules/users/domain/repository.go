package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Creator interface {
	Create(context.Context, *User) error
}
type ByEmailFinder interface {
	FindByEmail(context.Context, string) (*User, error)
}
type ByIDFinder interface {
	FindByID(context.Context, uuid.UUID) (*User, error)
}
type NameUpdater interface {
	UpdateName(context.Context, uuid.UUID, string) error
}
type Deleter interface {
	Delete(context.Context, uuid.UUID) error
}
type OwnershipChecker interface {
	HasOwnedCompanies(context.Context, uuid.UUID) (bool, error)
}
type VerificationTokenStore interface {
	FindByID(context.Context, uuid.UUID) (*User, error)
	SetVerificationToken(context.Context, uuid.UUID, string, time.Time) error
}
type EmailVerifier interface {
	VerifyEmail(context.Context, string) (*User, error)
}
