package app

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/douglas-franca/hireeasy/internal/platform/security/password"
)

type RegisterUserInput struct {
	Name     string
	Email    string
	Password string
}
type RegisterUser struct {
	creator domain.Creator
	pepper  string
	issuer  *VerificationTokenIssuer
}

func NewRegisterUser(creator domain.Creator, pepper string, issuer *VerificationTokenIssuer) *RegisterUser {
	return &RegisterUser{creator: creator, pepper: pepper, issuer: issuer}
}

func (u *RegisterUser) Execute(ctx context.Context, input RegisterUserInput) error {
	hash, err := password.Hash(input.Password, u.pepper)
	if err != nil {
		return err
	}

	user := domain.NewUser(domain.NewUserParams{Name: input.Name, Email: input.Email, PasswordHash: hash})
	if err := u.creator.Create(ctx, user); err != nil {
		return err
	}

	return u.issuer.Issue(ctx, user)
}
