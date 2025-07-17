package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url_shortener/internal/models"
	"url_shortener/internal/services"
)

var post_path string = "/api/shorten"

func TestPostOk(t *testing.T) {
	save_func := services.Saver
	services.Saver = func(data models.ToSave) error { return nil }
	defer func() {
		services.Saver = save_func
	}()
	req_body := `{"url": "any/url"}`
	req := httptest.NewRequest(http.MethodPost, post_path, strings.NewReader(req_body))
	w := httptest.NewRecorder()
	Post(w, req)
	response := w.Result()
	resp_body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201 got %d", response.StatusCode)
	}
	if !strings.HasPrefix(string(resp_body), `{"result":"http://localhost:8080/`) {
		t.Fatalf("expected body must starts from `http://localhost:8080/` but looks like `%s`", string(resp_body))
	}
}

func TestPostBadRequestBody(t *testing.T) {
	save_func := services.Saver
	services.Saver = func(data models.ToSave) error { return nil }
	defer func() {
		services.Saver = save_func
	}()
	req_body := `{"u": "any/url"}`
	req := httptest.NewRequest(http.MethodPost, post_path, strings.NewReader(req_body))
	w := httptest.NewRecorder()
	Post(w, req)
	response := w.Result()
	resp_body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400 got %d", response.StatusCode)
	}
	if string(resp_body) != "Bad request\n" {
		t.Fatalf("expected `Bad request` got %s", string(resp_body))
	}
}

func TestPostNotAllowed(t *testing.T) {
	save_func := services.Saver
	services.Saver = func(data models.ToSave) error { return nil }
	defer func() {
		services.Saver = save_func
	}()
	req := httptest.NewRequest(http.MethodDelete, post_path, nil)
	w := httptest.NewRecorder()
	Post(w, req)
	response := w.Result()
	resp_body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400 got %d", response.StatusCode)
	}
	if string(resp_body) != "method not allowed\n" {
		t.Fatalf("expected `Bad request` got %s", string(resp_body))
	}
}
