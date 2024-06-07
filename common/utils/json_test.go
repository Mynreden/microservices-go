package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"message": "hello world"}
	WriteJSON(rec, http.StatusOK, data)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var responseData map[string]string
	err := json.NewDecoder(rec.Body).Decode(&responseData)
	assert.NoError(t, err)
	assert.Equal(t, data, responseData)
}

func TestReadJSON(t *testing.T) {
	data := map[string]string{"message": "hello world"}
	body, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))

	var requestData map[string]string
	err := ReadJSON(req, &requestData)

	assert.NoError(t, err)
	assert.Equal(t, data, requestData)
}

func TestReadJSONError(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid json"))

	var requestData map[string]string
	err := ReadJSON(req, &requestData)

	assert.Error(t, err)
}

func TestWriteErr(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteErr(rec, http.StatusBadRequest, "bad request")

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var responseData map[string]string
	err := json.NewDecoder(rec.Body).Decode(&responseData)
	assert.NoError(t, err)
	assert.Equal(t, "bad request", responseData["error"])
}
