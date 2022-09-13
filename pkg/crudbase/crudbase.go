package crudbase

import (
	"fmt"
	"reflect"

	"github.com/go-chi/chi/v5"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

var Router chi.Mux

func Register(path string, m CrudBase) {
	if path[:1] != "/" {
		panic(fmt.Sprintf("path must start with \"/\", found: %v", path))
	}
	Router.Route(path, func(r chi.Router) {
		r.Post("/", HandleCreate(m))
		r.Get("/", HandleGetList)
		r.Get("/{id}", HandleGetDetail)
		r.Patch("/{id}", HandleUpdate)
		r.Delete("/{id}", HandleDelete)
	})

	mV := reflect.ValueOf(m)
	for mV.Kind() == reflect.Ptr {
		mV = mV.Elem()
	}
	a.AutoMigrate(mV.Interface())
}
