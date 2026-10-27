package app

import (
	"context"

	"github.com/douglas-franca/hionshop/internal/modules/users/domain"
	"github.com/douglas-franca/hionshop/internal/platform/security/password"
)

type Service struct {
	repo   domain.Repository
	pepper string
}

func NewService(repo domain.Repository, pepper string) *Service {
	return &Service{
		repo:   repo,
		pepper: pepper,
	}
}

func (s *Service) CreateUser(ctx context.Context, input CreateInput) error {
	hashedPassword, err := password.Hash(input.Password, s.pepper)
	if err != nil {
		return err
	}

	user := domain.NewUser(domain.NewUserParams{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPassword,
	})

	return s.repo.Create(ctx, user)
}
