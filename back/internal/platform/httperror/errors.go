package httperror

import (
	"errors"
	"log"
	"net/http"

	"github.com/douglas-franca/hireeasy/internal/platform/apperror"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Printf("panic recovered: %v", recovered)
				if !c.Writer.Written() {
					Write(c, apperror.Internal(errors.New("panic recovered")))
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}

type response struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}
		Write(c, c.Errors.Last().Err)
	}
}

func Abort(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Abort()
}

func Write(c *gin.Context, err error) {
	appErr := asApplicationError(err)
	c.JSON(statusFor(appErr.Category), response{Error: errorBody{Code: appErr.Code, Message: appErr.Message}})
}

func asApplicationError(err error) *apperror.Error {
	var appErr *apperror.Error
	if errors.As(err, &appErr) {
		return appErr
	}
	return apperror.Internal(err)
}

func statusFor(category apperror.Category) int {
	switch category {
	case apperror.CategoryValidation:
		return http.StatusBadRequest
	case apperror.CategoryUnauthenticated:
		return http.StatusUnauthorized
	case apperror.CategoryForbidden:
		return http.StatusForbidden
	case apperror.CategoryNotFound:
		return http.StatusNotFound
	case apperror.CategoryConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
