package service

import (
	"URLShorter/internal/storage"
	"URLShorter/internal/utils"
)

type URLService struct {
	store storage.URLStore
}

func NewURLService(store storage.URLStore) *URLService {
	return &URLService{store: store}
}

func (s *URLService) CreateShortURL(originalURL string) (string, error) {
	shortCode := utils.GenerateShortCode()
	err := s.store.Save(shortCode, originalURL)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	originalURL, err := s.store.Get(shortCode)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
