package test

import (
	"URLShorter/internal/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockStore struct {
	urls map[string]string
}

func (m *mockStore) SaveURL(originalURL, shortCode string) error {
	m.urls[shortCode] = originalURL
	return nil
}

func (m *mockStore) GetURL(shortCode string) (string, error) {
	url, exists := m.urls[shortCode]
	if !exists {
		return "", nil
	}
	return url, nil
}

func TestCreateShortURL(t *testing.T) {
	mock := &mockStore{urls: make(map[string]string)}
	service := service.NewURLService(mock)
	handler := NewURLHandler(service)

	request := map[string]string{
		"original_url": "https://www.example.com",
	}
	jsonData, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.CreateShortURL(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var response struct {
		ShortURL string `json:"short_url"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.ShortURL == "" {
		t.Error("expected non-empty short URL")
	}
}

func TestGetOriginalURL(t *testing.T) {
	mock := &mockStore{urls: make(map[string]string)}
	mock.urls["abc123"] = "https://www.example.com"
	service := service.NewURLService(mock)
	handler := NewURLHandler(service)

	req, err := http.NewRequest("GET", "/abc123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetOriginalURL(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var response struct {
		OriginalURL string `json:"original_url"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.OriginalURL != "https://www.example.com" {
		t.Errorf("expected URL 'https://www.example.com', got '%s'", response.OriginalURL)
	}
}
