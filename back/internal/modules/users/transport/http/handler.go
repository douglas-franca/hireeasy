package http

import (
	"context"
	nethttp "net/http"

	usersapp "github.com/douglas-franca/hionshop/internal/modules/users/app"
	"github.com/gin-gonic/gin"
)

type UserCreator interface {
	CreateUser(ctx context.Context, input usersapp.CreateInput) error
}

type Handler struct {
	users UserCreator
}

func NewHandler(users UserCreator) *Handler {
	return &Handler{
		users: users,
	}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.users.CreateUser(c.Request.Context(), usersapp.CreateInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{
			"error": "could not create user",
		})
		return
	}

	c.JSON(nethttp.StatusCreated, gin.H{
		"message": "user created",
	})
}
