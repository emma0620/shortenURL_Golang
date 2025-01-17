package handlers

import (
	"encoding/json"
	"net/http"
	"shortenURL_Golang/internal/services"
)

// HomeHandler 處理首頁
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the sURL shortener!"))
}

// ShortenHandler 處理短網址生成
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	shortURL := services.GenerateShortURL(req.URL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
	})
}

// RedirectHandler 處理短網址重定向
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	originalURL := services.GetOriginalURL("abc123")
	if originalURL == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
