package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainRequestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MainRequestHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("http status inesperado: retornou %v esperava %v",
			status, http.StatusOK)
	}

	expected := `Code.education Rocks!`
	if rr.Body.String() != expected {
		t.Errorf("retorno inesperado: retornou %v esperava %v",
			rr.Body.String(), expected)
	}
}
