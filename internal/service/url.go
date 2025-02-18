package service

import (
	"URLShortes/internal/repository"
	"context"
	"crypto/rand"
	"math/big"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const length = 10

type URLShortener interface {
	ShortenURL(ctx context.Context, originalURL string) (string, error)
	ExpandURL(ctx context.Context, shortURL string) (string, error)
}

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) ShortenURL(ctx context.Context, originalURL string) (string, error) {
	shortURL, err := generateShortURL()
	if err != nil {
		return "", err
	}
	err = s.repo.Save(ctx, shortURL, originalURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (s *URLService) ExpandURL(ctx context.Context, shortURL string) (string, error) {
	return s.repo.Get(ctx, shortURL)
}

func generateShortURL() (string, error) {
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		sb.WriteByte(charset[idx.Int64()])
	}
	return sb.String(), nil
}
