package routs

import (
	"net/http"
	"url_shortener/internal/handlers"
	"url_shortener/internal/middleware"
)

func GetRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/shorten", handlers.Post)
	mux.HandleFunc("/", handlers.Get)
	return middleware.Logger(mux)
}
