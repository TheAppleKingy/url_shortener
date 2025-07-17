package handlers

import (
	"net/http"
	"strconv"
	"strings"
	store "url_shortener/internal/storage"
)

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	store.Mu.RLock()
	var url string
	for _, v := range store.Storage {
		if id == v.Id {
			url = v.Old
		}
	}
	store.Mu.RUnlock()
	if url == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(url))
}
