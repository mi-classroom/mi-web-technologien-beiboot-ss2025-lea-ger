package tests

import (
	"backend/db"
	"backend/router"
	"bytes"
	"encoding/json"
	fmt "fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *gin.Engine {
	db.InitDB("host=localhost port=5432 user=admin password=admin dbname=iptc-editor sslmode=disable")
	return router.SetupRouter()
}

func uploadTestImage(server *gin.Engine, t *testing.T) int {
	if _, err := os.Stat("data/IPTC-PhotometadataRef-Std2024.1.jpg"); os.IsNotExist(err) {
		fmt.Println("Datei existiert nicht!")
	} else {
		fmt.Println("Datei gefunden.")
	}

	file, err := os.Open("data/IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)

	_, err = io.Copy(part, file)
	assert.NoError(t, err)
	writer.Close()

	req, err := http.NewRequest("POST", "/assets", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	respBody, _ := io.ReadAll(w.Body)
	err = json.Unmarshal(respBody, &response)
	assert.NoError(t, err)
	println(fmt.Printf("Response: %+v\n", response))

	id, ok := response["id"].(float64)
	assert.True(t, ok)

	println("ID:", int(id))
	return int(id)
}

func TestPingRoute(t *testing.T) {
	server := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetMetadata(t *testing.T) {
	server := setup()

	req, _ := http.NewRequest("GET", "/api/metadata/123", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestUploadAndDeleteImage(t *testing.T) {
	server := setup()
	id := uploadTestImage(server, t)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/assets/%d", id), nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPostMetadata(t *testing.T) {
	server := setup()
	id := uploadTestImage(server, t)

	req, _ := http.NewRequest("POST", "/api/metadata/-1", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

	req2, _ := http.NewRequest("POST", fmt.Sprintf("/api/metadata/%d", id), strings.NewReader(`{"Source": "KIKA Studios"}`))
	w2 := httptest.NewRecorder()
	server.ServeHTTP(w2, req2)

	println(w2.Body.String())

	assert.Equal(t, 200, w2.Code)

	req3, _ := http.NewRequest("GET", fmt.Sprintf("/api/metadata/%d", id), nil)
	w3 := httptest.NewRecorder()
	server.ServeHTTP(w3, req3)

	println(w3.Body.String())

	assert.Equal(t, 200, w3.Code)

	var response map[string]interface{}
	body, _ := io.ReadAll(w3.Body)
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, response["Source"], "KIKA Studios")
}

func TestUploadImageWithFolder(t *testing.T) {
	server := setup()

	folderName := "Testordner"
	folderBody := bytes.NewBufferString(`{"name": "` + folderName + `"}`)
	req, _ := http.NewRequest("POST", "/api/folders", folderBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var folderResp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &folderResp)
	assert.NoError(t, err)
	folderID := int(folderResp["id"].(float64))

	file, err := os.Open("data/IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err = writer.CreateFormFile("file", "IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)
	file.Seek(0, 0)
	io.Copy(body, file)
	writer.WriteField("folder_id", fmt.Sprintf("%d", folderID))
	writer.Close()

	req, err = http.NewRequest("POST", "/assets", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(folderID), response["folder_id"])
}

func TestUploadImageWithoutFolder(t *testing.T) {
	server := setup()
	file, err := os.Open("data/IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_, err = writer.CreateFormFile("file", "IPTC-PhotometadataRef-Std2024.1.jpg")
	assert.NoError(t, err)
	file.Seek(0, 0)
	io.Copy(body, file)
	writer.Close()

	req, err := http.NewRequest("POST", "/assets", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	_, hasFolder := response["folder_id"]
	assert.False(t, hasFolder)
}
