package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleProverbJson(t *testing.T) {
	// 1. HTTP Recorder erstellen
	recorder := httptest.NewRecorder()

	// 2. HTTP Request
	req, _ := http.NewRequest("GET", "/api", nil)

	// 3. HTTP Handler aufrufen
	HandleProverbJson(recorder, req)

	// 4. ERgebnis prüfen
	if recorder.Code != http.StatusOK {
		t.Errorf("Wrong status, got %v expected %v", recorder.Code, http.StatusOK)
	}
}
