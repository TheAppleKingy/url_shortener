package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"url_shortener/internal/models"
)

func LoadStorage(filename *string) {
	file, err := os.OpenFile(*filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		slog.Error("storage file error", "error", err)
		os.Exit(1)
	}
	defer file.Close()
	var data []models.ToSave
	var item models.ToSave
	decoder := json.NewDecoder(file)
	for {
		err := decoder.Decode(&item)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			slog.Error("parse storage error", "error", err)
			os.Exit(1)
		}
		data = append(data, item)
	}
	Mu.Lock()
	Storage = data
	if len(Storage) > 0 {
		Id = Storage[len(Storage)-1].Id + 1
	}
	Mu.Unlock()
}

func UpdateStorage(data models.ToSave) {
	Storage = append(Storage, data)
}
