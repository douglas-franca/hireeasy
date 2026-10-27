package postgres

import (
	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	usersqlc "github.com/douglas-franca/hireeasy/internal/modules/users/infra/postgres/sqlc"
	"github.com/google/uuid"
)

func toDomainUser(user *usersqlc.User) *domain.User {
	return &domain.User{
		ID:            uuid.UUID(user.ID.Bytes),
		Name:          user.Name,
		Email:         user.Email,
		PasswordHash:  user.PasswordHash,
		EmailVerified: user.EmailVerified,
	}
}
