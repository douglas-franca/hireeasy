package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
)

type ConfirmEmailVerification struct{ verifier domain.EmailVerifier }

func NewConfirmEmailVerification(verifier domain.EmailVerifier) *ConfirmEmailVerification {
	return &ConfirmEmailVerification{verifier: verifier}
}

func (u *ConfirmEmailVerification) Execute(ctx context.Context, token string) error {
	_, err := u.verifier.VerifyEmail(ctx, hashVerificationToken(token))

	return err
}
