package http

import "github.com/google/uuid"

type createUserResponse struct {
	Message string `json:"message"`
}
type currentUserResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
}
