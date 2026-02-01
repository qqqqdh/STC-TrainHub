package http

import (
	"net/http"

	"stc-trainhub/internal/config"
)

func NewRouter(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("GET /health", handleHealth)

	// middlewares
	var h http.Handler = mux
	h = RecoveryMiddleware(h)
	h = RequestLogMiddleware(h)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		OK(w, map[string]any{
			"service": "STC-TrainHub API",
			"version": "v0.1",
		})
	})

	return h
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	OK(w, map[string]any{
		"status": "ok",
	})
}
