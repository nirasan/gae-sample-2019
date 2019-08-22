package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			res.Code,
			http.StatusOK,
		)
	}

	expected := "Hello World"
	if res.Body.String() != expected {
		t.Errorf(
			"unexpected body: got (%v) want (%v)",
			res.Body.String(),
			expected,
		)
	}
}
