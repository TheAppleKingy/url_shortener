package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"url_shortener/internal/routs"
	"url_shortener/internal/storage"
)

func main() {
	filename := flag.String("f", "", "Parse filename of file for urls")
	flag.Parse()
	storage.LoadStorage(filename)
	router := routs.GetRouter()
	slog.Info("Starting server. Port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		slog.Error("server fail", "error", err)
		os.Exit(1)
	}
}
