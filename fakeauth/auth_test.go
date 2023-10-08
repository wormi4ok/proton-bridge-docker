package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Auth-Method", "plain")
	req.Header.Set("Auth-User", "user")
	req.Header.Set("Auth-Pass", "password")
	req.Header.Set("Auth-Protocol", "imap")
	req.Header.Set("Auth-Login-Attempt", "1")

	rr := httptest.NewRecorder()
	handler := AuthHandler(config)

	handler.ServeHTTP(rr, req)

	var (
		statusCode, expectedStatusCode = rr.Code, http.StatusOK
		status, expectedStatus         = rr.Header().Get("Auth-Status"), "OK"
		authPort, expectedAuthPort     = rr.Header().Get("Auth-Port"), "1143"
	)
	if statusCode != expectedStatusCode {
		t.Fatalf("Unexpected response status code: %d", statusCode)
	}
	if status != expectedStatus {
		t.Fatalf("Unexpected status in response: %s", status)
	}

	if authPort != expectedAuthPort {
		t.Errorf("Unexpected auth port in response: %s", authPort)
	}
}
