package storage

type URLStore interface {
	Save(shortCode, originalURL string) error
	Get(shortCode string) (string, error)
}
