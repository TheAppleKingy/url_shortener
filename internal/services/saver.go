package services

import (
	"encoding/json"
	"os"
	"url_shortener/internal/models"
	"url_shortener/internal/storage"
)

var Saver = saver

func saver(data models.ToSave) error {
	file, err := os.OpenFile("urls.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	storage.UpdateStorage(data)
	return nil
}
