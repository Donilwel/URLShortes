package service_test

import (
	"context"
	"testing"

	"URLShortes/internal/repository"
)

func TestSaveAndGetURL(t *testing.T) {
	repo := repository.NewInMemoryRepository()

	originalURL := "https://example.com"
	shortURL := "A1b2C3d4E5"

	err := repo.Save(context.Background(), shortURL, originalURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	retrievedURL, err := repo.Get(context.Background(), shortURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if retrievedURL != originalURL {
		t.Errorf("expected %s, got %s", originalURL, retrievedURL)
	}
}

func TestGetNonExistentURL(t *testing.T) {
	repo := repository.NewInMemoryRepository()

	_, err := repo.Get(context.Background(), "NonExistent123")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
