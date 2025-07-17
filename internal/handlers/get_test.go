package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_shortener/internal/models"
	"url_shortener/internal/storage"
)

var path = "/%d"

func setup() {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	storage.Storage = []models.ToSave{
		{Id: 1, Old: "https://example.com", New: "abc123"},
	}
}

func TestGetOk(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf(path, 1), nil)
	w := httptest.NewRecorder()
	Get(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusTemporaryRedirect {
		t.Fatalf("expected status 307, got %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "https://example.com" {
		t.Fatalf("expected `https://example.com` got %s", string(body))
	}
}

func TestGetNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf(path, 2), nil)
	w := httptest.NewRecorder()
	Get(w, req)
	response := w.Result()
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "Not found\n" {
		t.Fatalf("expected `Not found` got %s", string(body))
	}
}

func TestNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf(path, 2), nil)
	w := httptest.NewRecorder()
	Get(w, req)
	response := w.Result()
	body, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 405, got %d", response.StatusCode)
	}
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "method not allowed\n" {
		t.Fatalf("expected `method not allowed` got %s", string(body))
	}
}

func init() {
	setup()
}
