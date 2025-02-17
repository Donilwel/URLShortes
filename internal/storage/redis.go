package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(address string) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	// Пытаемся подключиться к Redis
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{client: client}, nil
}

func (r *RedisStore) SaveURL(originalURL, shortCode string) error {
	err := r.client.Set(context.Background(), shortCode, originalURL, 0).Err()
	if err != nil {
		log.Printf("Failed to save URL: %v", err)
		return err
	}
	return nil
}

func (r *RedisStore) GetURL(shortCode string) (string, error) {
	val, err := r.client.Get(context.Background(), shortCode).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		log.Printf("Failed to retrieve URL: %v", err)
		return "", err
	}

	return val, nil
}
