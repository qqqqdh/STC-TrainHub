package app

import (
	"net/http"

	"stc-trainhub/internal/config"
	stchttp "stc-trainhub/internal/http"
)

type App struct {
	cfg     config.Config
	handler http.Handler
}

func New(cfg config.Config) (*App, error) {
	r := stchttp.NewRouter(cfg)
	return &App{
		cfg:     cfg,
		handler: r,
	}, nil
}

func (a *App) Handler() http.Handler {
	return a.handler
}
