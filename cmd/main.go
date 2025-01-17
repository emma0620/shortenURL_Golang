package main

import (
	"log"
	"net/http"
	"shortenURL_Golang/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// 路由設定
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/shorten", handlers.ShortenHandler)
	mux.HandleFunc("/abc123", handlers.RedirectHandler)

	// 啟動伺服器
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
