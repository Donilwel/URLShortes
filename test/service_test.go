package service_test

import (
	"context"
	"testing"

	"URLShortes/internal/repository"
	"URLShortes/internal/service"
)

func TestExpandURL(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	srv := service.NewURLService(repo)

	originalURL := "https://example.com"
	shortURL, _ := srv.ShortenURL(context.Background(), originalURL)

	retrievedURL, err := srv.ExpandURL(context.Background(), shortURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if retrievedURL != originalURL {
		t.Errorf("expected %s, got %s", originalURL, retrievedURL)
	}
}

func TestExpandNonExistentURL(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	srv := service.NewURLService(repo)

	_, err := srv.ExpandURL(context.Background(), "NonExist123")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
