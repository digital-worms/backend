package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type healthResponse struct {
	Status string `json:"status"`
}

func TestHealth(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	router := NewRouter(&stubDatabasePinger{})
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected HTTP status: got %d, want %d", recorder.Code, http.StatusOK)
	}
	var response healthResponse
	const expectedStatus = "ok"

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != expectedStatus {
		t.Errorf("unexpected response status: got %q, want %q", response.Status, expectedStatus)
	}

	const expectedContentType = "application/json; charset=utf-8"

	contentType := recorder.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("unexpected Content-Type: got %q, want %q", contentType, expectedContentType)
	}

}

func TestReadySuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	router := NewRouter(&stubDatabasePinger{})
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected HTTP status: got %d, want %d", recorder.Code, http.StatusOK)
		return
	}
	var response healthResponse
	const expectedStatus = "ok"

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != expectedStatus {
		t.Errorf("unexpected response status: got %q, want %q", response.Status, expectedStatus)
	}

	const expectedContentType = "application/json; charset=utf-8"

	contentType := recorder.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("unexpected Content-Type: got %q, want %q", contentType, expectedContentType)
	}

}

var errReadyUnavailable = errors.New("database unavailable")

func TestReadyUnavailable(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	router := NewRouter(&stubDatabasePinger{
		pingErr: errReadyUnavailable,
	})
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusServiceUnavailable {
		t.Errorf("unexpected HTTP status: got %d, want %d", recorder.Code, http.StatusServiceUnavailable)
		return
	}
	var response healthResponse
	const expectedStatus = "unavailable"

	if strings.Contains(recorder.Body.String(), errReadyUnavailable.Error()) {
		t.Error("internal database error leaked in the HTTP response")
	}

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	var raw json.RawMessage
	err = json.NewDecoder(recorder.Body).Decode(&raw)
	if err != io.EOF {
		t.Errorf("expected end of response body (io.EOF), got: %v", err)
	}

	if response.Status != expectedStatus {
		t.Errorf("unexpected response status: got %q, want %q", response.Status, expectedStatus)
	}

	const expectedContentType = "application/json; charset=utf-8"

	contentType := recorder.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("unexpected Content-Type: got %q, want %q", contentType, expectedContentType)
	}
}
