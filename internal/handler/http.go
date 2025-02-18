package handler

import (
	"URLShortes/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type URLHandler struct {
	service service.URLShortener
}

func NewURLHandler(s service.URLShortener) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", h.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", h.ExpandURL).Methods("GET")
	return r
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	short, err := h.service.ShortenURL(r.Context(), req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"short_url": short})
}

func (h *URLHandler) ExpandURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	original, err := h.service.ExpandURL(r.Context(), vars["shortURL"])
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
}
