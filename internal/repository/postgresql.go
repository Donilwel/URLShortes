package repository

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(ctx context.Context, shortURL, originalURL string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO urls (short_url, original_url) VALUES ($1, $2) ON CONFLICT (original_url) DO NOTHING", shortURL, originalURL)
	return err
}

func (r *PostgresRepository) Get(ctx context.Context, shortURL string) (string, error) {
	var originalURL string
	err := r.db.QueryRowContext(ctx, "SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	return originalURL, err
}
