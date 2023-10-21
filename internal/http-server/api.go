package httpserver

import (
	"emobletest/internal/storage"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

type API struct {
	db     storage.StorageInterface
	r      *chi.Mux
	logger *slog.Logger
}

func New(db storage.StorageInterface, log *slog.Logger) *API {
	a := API{db: db, r: chi.NewRouter(), logger: log}
	a.endpoints()
	return &a
}

func (api *API) Router() *chi.Mux {
	return api.r
}
