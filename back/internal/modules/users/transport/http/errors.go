package http

import "github.com/douglas-franca/hireeasy/internal/platform/apperror"

func validationError() *apperror.Error {
	return apperror.New(apperror.CodeValidationError, apperror.CategoryValidation, "The request is invalid.")
}
