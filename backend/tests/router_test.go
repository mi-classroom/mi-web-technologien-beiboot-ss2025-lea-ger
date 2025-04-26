package tests

import (
	"backend/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	server := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetMetadata(t *testing.T) {
	server := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/api/metadata/123", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message": "read metadata", "id": "123"}`, w.Body.String())
}

func TestPostMetadata(t *testing.T) {
	server := router.SetupRouter()

	req, _ := http.NewRequest("POST", "/api/metadata/123", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message": "write metadata", "id": "123"}`, w.Body.String())
}
