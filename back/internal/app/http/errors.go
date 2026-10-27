package http

import "github.com/douglas-franca/hireeasy/internal/platform/apperror"

func httperrorNotFound() *apperror.Error {
	return apperror.New("ROUTE_NOT_FOUND", apperror.CategoryNotFound, "Route not found.")
}
