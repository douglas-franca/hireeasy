package apperror

import "fmt"

type Category string

const (
	CategoryValidation      Category = "validation"
	CategoryUnauthenticated Category = "unauthenticated"
	CategoryForbidden       Category = "forbidden"
	CategoryNotFound        Category = "not_found"
	CategoryConflict        Category = "conflict"
	CategoryInternal        Category = "internal"
)

const (
	CodeValidationError          = "VALIDATION_ERROR"
	CodeInternalError            = "INTERNAL_ERROR"
	CodeUnauthenticated          = "UNAUTHENTICATED"
	CodeInvalidCredentials       = "INVALID_CREDENTIALS"
	CodeEmailNotVerified         = "EMAIL_NOT_VERIFIED"
	CodeVerificationTokenInvalid = "VERIFICATION_TOKEN_INVALID"
	CodeEmailAlreadyExists       = "EMAIL_ALREADY_EXISTS"
	CodeCompanyNotFound          = "COMPANY_NOT_FOUND"
	CodeUserOwnsCompany          = "USER_OWNS_COMPANY"
)

type Error struct {
	Code     string
	Category Category
	Message  string
	Cause    error
}

func (e *Error) Error() string { return e.Code }
func (e *Error) Unwrap() error { return e.Cause }

func New(code string, category Category, message string) *Error {
	return &Error{Code: code, Category: category, Message: message}
}

func Wrap(base *Error, cause error) *Error {
	return &Error{Code: base.Code, Category: base.Category, Message: base.Message, Cause: cause}
}

func Internal(cause error) *Error {
	return &Error{Code: CodeInternalError, Category: CategoryInternal, Message: "An unexpected error occurred.", Cause: fmt.Errorf("%w", cause)}
}
