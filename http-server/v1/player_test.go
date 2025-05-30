package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/checks", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}
	resp := httptest.NewRecorder()

	PlayerServer(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
	}
	if resp.Body.String() != "Player Server is running" {
		t.Errorf("Expected response body 'Player Server is running', got '%s'", resp.Body.String())
	}
}
