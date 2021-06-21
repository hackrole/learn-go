package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelathCheckhandler ...
func TestHelathCheckhandler(r *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandlern)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler rturn body: Get %v want %v", rr.Body.String(), expected)
	}
}
