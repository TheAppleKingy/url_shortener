package handlers

import (
	"encoding/json"
	"net/http"
	"url_shortener/internal/models"
	"url_shortener/internal/services"
	store "url_shortener/internal/storage"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var data models.ReqData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if data.Url == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	store.Mu.Lock()
	new_url := services.URLShortener()
	record := models.ToSave{
		Id: store.Id, Old: data.Url, New: new_url,
	}
	store.Id++
	store.Mu.Unlock()

	err := services.Saver(record)
	if err != nil {
		http.Error(w, "storage error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := models.RespData{Result: new_url}
	json.NewEncoder(w).Encode(response)
}
