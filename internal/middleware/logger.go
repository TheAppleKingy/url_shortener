package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWrapper struct {
	http.ResponseWriter
	statusCode int
	rSize      int
}

func (rw *responseWrapper) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)
	rw.statusCode = code
}

func (rw *responseWrapper) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.rSize = size
	return size, err
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		respWriter := &responseWrapper{
			ResponseWriter: w,
			statusCode:     400,
			rSize:          0,
		}
		handler.ServeHTTP(respWriter, r)
		end := time.Since(start)
		slog.Info(
			"REQUEST",
			"METHOD", r.Method,
			"URL", r.URL.Path,
			"DURATION", end,
		)
		slog.Info(
			"RESPONSE",
			"URL", r.URL.Path,
			"STATUS", respWriter.statusCode,
			"SIZE", respWriter.rSize,
		)
	})
}
