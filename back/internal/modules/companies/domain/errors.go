package domain

import "github.com/douglas-franca/hireeasy/internal/platform/apperror"

var ErrEmailNotVerified = apperror.New(apperror.CodeEmailNotVerified, apperror.CategoryForbidden, "Verify your email before continuing.")
var ErrCompanyNotFound = apperror.New(apperror.CodeCompanyNotFound, apperror.CategoryNotFound, "Company not found.")
