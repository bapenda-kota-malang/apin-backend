package routerhelper

import (
	"github.com/go-chi/chi/v5"
)

func RegCrud(r *chi.Mux, path string, c CrudBase) {
	r.Route(path, func(r chi.Router) {
		r.Post("/", c.Create)
		r.Get("/", c.GetList)
		r.Get("/{id}", c.GetDetail)
		r.Patch("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})
}
