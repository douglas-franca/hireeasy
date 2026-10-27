package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSwaggerDocumentIsServed(t *testing.T) {
	router := NewRouter()
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/swagger/openapi.json", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("got %d", response.Code)
	}
	if response.Header().Get("Content-Type") != "application/json" {
		t.Fatal("expected JSON OpenAPI document")
	}
}
