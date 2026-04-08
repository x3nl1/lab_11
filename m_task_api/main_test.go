package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest("POST", "/calculate", bytes.NewBuffer([]byte(`{"values":[1,2,3]}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected 200, got %d", w.Code)
	}

	expected := `{"sum":6}`
	if w.Body.String() != expected {
		t.Errorf("expected %s, got %s", expected, w.Body.String())
	}
}

func TestCalculateInvalid(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest("POST", "/calculate", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestProtectedNoToken(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 401 {
		t.Errorf("expected 401, got %d", w.Code)
	}
}