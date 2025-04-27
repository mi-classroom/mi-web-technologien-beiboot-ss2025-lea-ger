package tests

import (
	"backend/router"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	strings "strings"
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

	assert.Equal(t, 404, w.Code)

	req2, _ := http.NewRequest("GET", "/api/metadata/1", nil)
	w2 := httptest.NewRecorder()
	server.ServeHTTP(w2, req2)

	println(w2.Body.String())

	assert.Equal(t, 200, w2.Code)

	var response map[string]interface{}
	body, _ := io.ReadAll(w2.Body)
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, response["ImageHeight"], 1280.)
	assert.Equal(t, response["ImageWidth"], 853.)
	assert.Equal(t, response["FileType"], "JPEG")
}

func TestPostMetadata(t *testing.T) {
	server := router.SetupRouter()

	req, _ := http.NewRequest("POST", "/api/metadata/123", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

	req2, _ := http.NewRequest("POST", "/api/metadata/1", strings.NewReader(`{"Author": "KIKA Studios"}`))
	w2 := httptest.NewRecorder()
	server.ServeHTTP(w2, req2)

	println(w2.Body.String())

	assert.Equal(t, 200, w2.Code)

	req3, _ := http.NewRequest("GET", "/api/metadata/1", nil)
	w3 := httptest.NewRecorder()
	server.ServeHTTP(w3, req3)

	println(w3.Body.String())

	assert.Equal(t, 200, w3.Code)

	var response map[string]interface{}
	body, _ := io.ReadAll(w3.Body)
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, response["Author"], "KIKA Studios")
}
