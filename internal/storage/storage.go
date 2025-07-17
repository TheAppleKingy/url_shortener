package storage

import (
	"sync"
	"url_shortener/internal/models"
)

var Id int = 1
var Mu sync.RWMutex
var Storage []models.ToSave
