package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/google/uuid"
)

type RequestEmailVerification struct {
	finder domain.ByIDFinder
	issuer *VerificationTokenIssuer
}

func NewRequestEmailVerification(finder domain.ByIDFinder, issuer *VerificationTokenIssuer) *RequestEmailVerification {
	return &RequestEmailVerification{finder: finder, issuer: issuer}
}

func (u *RequestEmailVerification) Execute(ctx context.Context, userID uuid.UUID) error {
	user, err := u.finder.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.EmailVerified {
		return nil
	}

	return u.issuer.Issue(ctx, user)
}
