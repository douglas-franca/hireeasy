package http

import (
	"context"
	companiesapp "github.com/douglas-franca/hireeasy/internal/modules/companies/app"
	"github.com/douglas-franca/hireeasy/internal/modules/companies/domain"
	sessionshttp "github.com/douglas-franca/hireeasy/internal/modules/sessions/transport/http"
	"github.com/douglas-franca/hireeasy/internal/platform/apperror"
	"github.com/douglas-franca/hireeasy/internal/platform/httperror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	nethttp "net/http"
)

type createCompanyUseCase interface {
	Execute(context.Context, uuid.UUID, companiesapp.CreateCompanyInput) (*domain.Company, error)
}
type listCompaniesUseCase interface {
	Execute(context.Context, uuid.UUID) ([]domain.Company, error)
}
type getCompanyUseCase interface {
	Execute(context.Context, uuid.UUID, uuid.UUID) (*domain.Company, error)
}
type Handler struct {
	create createCompanyUseCase
	list   listCompaniesUseCase
	get    getCompanyUseCase
}

func NewHandler(create createCompanyUseCase, list listCompaniesUseCase, get getCompanyUseCase) *Handler {
	return &Handler{create: create, list: list, get: get}
}
func (h *Handler) Create(c *gin.Context) {
	var req createCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httperror.Abort(c, validationError())
		return
	}
	company, err := h.create.Execute(c.Request.Context(), sessionshttp.UserID(c), companiesapp.CreateCompanyInput{Name: req.Name})
	if err != nil {
		httperror.Abort(c, err)
		return
	}
	c.JSON(nethttp.StatusCreated, companyResponseFrom(company))
}
func (h *Handler) List(c *gin.Context) {
	companies, err := h.list.Execute(c.Request.Context(), sessionshttp.UserID(c))
	if err != nil {
		httperror.Abort(c, err)
		return
	}
	c.JSON(nethttp.StatusOK, companyResponsesFrom(companies))
}
func (h *Handler) Get(c *gin.Context) {
	companyID, err := uuid.Parse(c.Param("companyID"))
	if err != nil {
		httperror.Abort(c, validationError())
		return
	}
	company, err := h.get.Execute(c.Request.Context(), sessionshttp.UserID(c), companyID)
	if err != nil {
		httperror.Abort(c, err)
		return
	}
	c.JSON(nethttp.StatusOK, companyResponseFrom(company))
}
func validationError() *apperror.Error {
	return apperror.New(apperror.CodeValidationError, apperror.CategoryValidation, "The request is invalid.")
}
