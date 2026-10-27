package domain

import (
	"github.com/google/uuid"
)

type NewUserParams struct {
	Name         string
	Email        string
	PasswordHash string
}

type User struct {
	ID            uuid.UUID
	Name          string
	Email         string
	PasswordHash  string
	EmailVerified bool
}

func NewUser(params NewUserParams) *User {
	return &User{
		ID:            uuid.Must(uuid.NewV7()),
		Name:          params.Name,
		Email:         params.Email,
		PasswordHash:  params.PasswordHash,
		EmailVerified: false,
	}
}
