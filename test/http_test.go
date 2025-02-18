package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"URLShortes/internal/handler"
	"URLShortes/internal/repository"
	"URLShortes/internal/service"
)

func TestShortenURLHandler(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	srv := service.NewURLService(repo)
	h := handler.NewURLHandler(srv)

	server := httptest.NewServer(h.Router())
	defer server.Close()

	reqBody := `{"url": "https://example.com"}`
	resp, err := http.Post(server.URL+"/shorten", "application/json", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)

	shortURL, ok := response["short_url"]
	if !ok || len(shortURL) != 10 {
		t.Errorf("expected valid short_url, got %v", response)
	}
}

func TestExpandNonExistentURLHandler(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	srv := service.NewURLService(repo)
	h := handler.NewURLHandler(srv)

	server := httptest.NewServer(h.Router())
	defer server.Close()

	resp, err := http.Get(server.URL + "/NonExistent123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", resp.StatusCode)
	}
}
