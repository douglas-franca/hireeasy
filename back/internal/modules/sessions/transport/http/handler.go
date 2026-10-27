package http

import (
	"context"
	sessionsapp "github.com/douglas-franca/hireeasy/internal/modules/sessions/app"
	"github.com/douglas-franca/hireeasy/internal/platform/apperror"
	"github.com/douglas-franca/hireeasy/internal/platform/httperror"
	"github.com/gin-gonic/gin"
	nethttp "net/http"
	"time"
)

type createSessionUseCase interface {
	Execute(context.Context, sessionsapp.CreateSessionInput) (string, error)
}
type deleteSessionUseCase interface {
	Execute(context.Context, string) error
}
type Handler struct {
	create   createSessionUseCase
	delete   deleteSessionUseCase
	secure   bool
	duration time.Duration
}

func NewHandler(create createSessionUseCase, delete deleteSessionUseCase, secure bool, duration time.Duration) *Handler {
	return &Handler{create: create, delete: delete, secure: secure, duration: duration}
}
func (h *Handler) Create(c *gin.Context) {
	var req createSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperror.Abort(c, validationError())
		return
	}
	token, err := h.create.Execute(c.Request.Context(), sessionsapp.CreateSessionInput{Email: req.Email, Password: req.Password})
	if err != nil {
		httperror.Abort(c, err)
		return
	}
	c.SetCookie(CookieName, token, int(h.duration.Seconds()), "/", "", h.secure, true)
	c.Status(nethttp.StatusNoContent)
}
func (h *Handler) Delete(c *gin.Context) {
	token, _ := c.Cookie(CookieName)
	if err := h.delete.Execute(c.Request.Context(), token); err != nil {
		httperror.Abort(c, err)
		return
	}
	c.SetCookie(CookieName, "", -1, "/", "", h.secure, true)
	c.Status(nethttp.StatusNoContent)
}
func validationError() *apperror.Error {
	return apperror.New(apperror.CodeValidationError, apperror.CategoryValidation, "The request is invalid.")
}
