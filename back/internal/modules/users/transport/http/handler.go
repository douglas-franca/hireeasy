package http

import (
	"context"
	nethttp "net/http"

	sessionshttp "github.com/douglas-franca/hireeasy/internal/modules/sessions/transport/http"
	usersapp "github.com/douglas-franca/hireeasy/internal/modules/users/app"
	"github.com/douglas-franca/hireeasy/internal/modules/users/domain"
	"github.com/douglas-franca/hireeasy/internal/platform/httperror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type registerUserUseCase interface {
	Execute(context.Context, usersapp.RegisterUserInput) error
}
type getCurrentUserUseCase interface {
	Execute(context.Context, uuid.UUID) (*domain.User, error)
}
type updateCurrentUserUseCase interface {
	Execute(context.Context, uuid.UUID, usersapp.UpdateCurrentUserInput) error
}
type deleteCurrentUserUseCase interface {
	Execute(context.Context, uuid.UUID) error
}
type requestVerificationUseCase interface {
	Execute(context.Context, uuid.UUID) error
}
type confirmVerificationUseCase interface {
	Execute(context.Context, string) error
}
type Handler struct {
	register            registerUserUseCase
	get                 getCurrentUserUseCase
	update              updateCurrentUserUseCase
	delete              deleteCurrentUserUseCase
	requestVerification requestVerificationUseCase
	confirmVerification confirmVerificationUseCase
}

func NewHandler(
	register registerUserUseCase,
	get getCurrentUserUseCase,
	update updateCurrentUserUseCase,
	delete deleteCurrentUserUseCase,
	requestVerification requestVerificationUseCase,
	confirmVerification confirmVerificationUseCase,
) *Handler {
	return &Handler{
		register:            register,
		get:                 get,
		update:              update,
		delete:              delete,
		requestVerification: requestVerification,
		confirmVerification: confirmVerification,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req registerUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperror.Abort(c, validationError())
		return
	}

	input := usersapp.RegisterUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.register.Execute(
		c.Request.Context(),
		input,
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.JSON(nethttp.StatusCreated, createUserResponse{
		Message: "User created.",
	})
}

func (h *Handler) GetCurrent(c *gin.Context) {
	user, err := h.get.Execute(
		c.Request.Context(),
		sessionshttp.UserID(c),
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.JSON(nethttp.StatusOK, currentUserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
	})
}

func (h *Handler) UpdateCurrent(c *gin.Context) {
	var req updateCurrentUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperror.Abort(c, validationError())
		return
	}

	input := usersapp.UpdateCurrentUserInput{
		Name: req.Name,
	}

	err := h.update.Execute(
		c.Request.Context(),
		sessionshttp.UserID(c),
		input,
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.Status(nethttp.StatusNoContent)
}

func (h *Handler) DeleteCurrent(c *gin.Context) {
	err := h.delete.Execute(
		c.Request.Context(),
		sessionshttp.UserID(c),
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.Status(nethttp.StatusNoContent)
}

func (h *Handler) RequestVerification(c *gin.Context) {
	err := h.requestVerification.Execute(
		c.Request.Context(),
		sessionshttp.UserID(c),
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.Status(nethttp.StatusNoContent)
}

func (h *Handler) ConfirmVerification(c *gin.Context) {
	var req confirmEmailVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperror.Abort(c, validationError())
		return
	}

	err := h.confirmVerification.Execute(
		c.Request.Context(),
		req.Token,
	)
	if err != nil {
		httperror.Abort(c, err)
		return
	}

	c.Status(nethttp.StatusNoContent)
}
