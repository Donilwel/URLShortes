package main

import (
	"URLShortes/internal/config"
	"URLShortes/internal/handler"
	"URLShortes/internal/repository"
	"URLShortes/internal/service"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	var repo repository.URLRepository
	if cfg.StorageType == "postgres" {
		db, err := sql.Open("postgres", cfg.PostgresDSN)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		repo = repository.NewPostgresRepository(db)
	} else {
		repo = repository.NewInMemoryRepository()
	}

	urlService := service.NewURLService(repo)
	httpHandler := handler.NewURLHandler(urlService)

	log.Println("Starting server on port", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, httpHandler.Router())
}
