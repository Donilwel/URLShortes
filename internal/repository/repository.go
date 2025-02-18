package repository

import "context"

type URLRepository interface {
	Save(ctx context.Context, shortURL, originalURL string) error
	Get(ctx context.Context, shortURL string) (string, error)
}
