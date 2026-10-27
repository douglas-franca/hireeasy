package httperror

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglas-franca/hireeasy/internal/platform/apperror"
	"github.com/gin-gonic/gin"
)

func TestMiddlewareWritesKeyedErrorContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Middleware())
	router.GET("/", func(c *gin.Context) {
		Abort(c, apperror.New(apperror.CodeEmailNotVerified, apperror.CategoryForbidden, "Verify your email before continuing."))
	})
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/", nil))
	if response.Code != http.StatusForbidden {
		t.Fatalf("got %d", response.Code)
	}
	expected := `{"error":{"code":"EMAIL_NOT_VERIFIED","message":"Verify your email before continuing."}}`
	if response.Body.String() != expected {
		t.Fatalf("got %s", response.Body.String())
	}
}
