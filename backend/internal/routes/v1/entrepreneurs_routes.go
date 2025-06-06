package v1

import (
	"github.com/cyarleque/sumaq/internal/drivers"
	"github.com/go-chi/chi/v5"
)

func (config ConnectionConfig) EntrepreneursRouter() chi.Router {
	r := chi.NewRouter()
	driver := drivers.NewEntrepreneursDriver(config.SumaqDB)

	r.Get("/", driver.GetEntrepreneurs)
	r.Put("/{id}", driver.UpdateUrlEntrepreneur)
	r.Get("/{id}", driver.GetEntrepreneur)

	return r
}
