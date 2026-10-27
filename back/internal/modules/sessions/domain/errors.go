package domain

import "github.com/douglas-franca/hireeasy/internal/platform/apperror"

var ErrInvalidCredentials = apperror.New(apperror.CodeInvalidCredentials, apperror.CategoryUnauthenticated, "Invalid email or password.")
var ErrUnauthenticated = apperror.New(apperror.CodeUnauthenticated, apperror.CategoryUnauthenticated, "Authentication is required.")
