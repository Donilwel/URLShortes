package handler

import (
	"URLShorter/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		OriginalURL string `json:"original_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortCode, err := service.CreateShortURL(request.OriginalURL)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		return
	}

	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: "http://localhost:8080/" + shortCode,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]
	originalURL, err := h.service.GetOriginalURL(shortCode)

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
