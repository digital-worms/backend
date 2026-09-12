package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type healthResponse struct {
	Status string `json:"status"`
}

func TestHealth(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	router := NewRouter()
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("Ожидали: %d получили: %d", http.StatusOK, recorder.Code)
	}

	var response healthResponse
	const expectedStatus = "ok"

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != expectedStatus {
		t.Errorf("Ожидали: %s получили: %s", expectedStatus, response.Status)
	}

	const expectedContentType = "application/json; charset=utf-8"

	contentType := recorder.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("ожидался: %s, получили: %s", expectedContentType, contentType)
	}

}
