package domain

import "github.com/douglas-franca/hireeasy/internal/platform/apperror"

var ErrEmailAlreadyExists = apperror.New(apperror.CodeEmailAlreadyExists, apperror.CategoryConflict, "An account with this email already exists.")
var ErrVerificationTokenInvalid = apperror.New(apperror.CodeVerificationTokenInvalid, apperror.CategoryValidation, "The verification token is invalid or expired.")
var ErrUserOwnsCompany = apperror.New(apperror.CodeUserOwnsCompany, apperror.CategoryConflict, "Transfer or delete owned companies before deleting your account.")
