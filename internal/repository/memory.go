package repository

import (
	"context"
	"errors"
	"sync"
)

type InMemoryRepository struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[string]string)}
}

func (r *InMemoryRepository) Save(ctx context.Context, shortURL, originalURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.data[originalURL]; exists {
		return errors.New("URL already exists")
	}
	r.data[shortURL] = originalURL
	return nil
}

func (r *InMemoryRepository) Get(ctx context.Context, shortURL string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, exists := r.data[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}
	return url, nil
}
