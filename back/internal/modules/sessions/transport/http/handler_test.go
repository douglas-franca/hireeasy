package http

import (
	"context"
	sessionsapp "github.com/douglas-franca/hireeasy/internal/modules/sessions/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type createSessionStub struct{}

func (createSessionStub) Execute(context.Context, sessionsapp.CreateSessionInput) (string, error) {
	return "token", nil
}

type deleteSessionStub struct{}

func (deleteSessionStub) Execute(context.Context, string) error { return nil }
func TestCreateSessionSetsHTTPOnlyCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/sessions", NewHandler(createSessionStub{}, deleteSessionStub{}, false, time.Hour).Create)
	request := httptest.NewRequest(http.MethodPost, "/sessions", strings.NewReader(`{"email":"ana@example.com","password":"password123"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusNoContent {
		t.Fatalf("got %d", response.Code)
	}
	cookie := response.Result().Cookies()[0]
	if cookie.Name != CookieName || !cookie.HttpOnly {
		t.Fatal("expected an HttpOnly session cookie")
	}
}
