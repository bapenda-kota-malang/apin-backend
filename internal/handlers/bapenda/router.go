package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/ppat"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		r.Patch("/reset-password", account.ResetPassword)
		r.Patch("/change-password", account.ChangePassword)
	})

	r.Route("/pegawai", func(r chi.Router) {
		r.Post("/", pegawai.Create)
		r.Get("/", pegawai.GetList)
		r.Get("/{id}", pegawai.GetDetail)
		r.Patch("/{id}", pegawai.Update)
		r.Delete("/{id}", pegawai.Delete)
	})

	r.Route("/ppat", func(r chi.Router) {
		r.Post("/", ppat.Create)
		r.Get("/", ppat.GetList)
		r.Get("/{id}", ppat.GetDetail)
		r.Patch("/{id}", ppat.Update)
		r.Delete("/{id}", ppat.Delete)
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", user.Create)
		r.Get("/", user.GetList)
		r.Get("/{id}", user.GetDetail)
		r.Patch("/{id}", user.Update)
		r.Delete("/{id}", user.Delete)
	})

	r.Route("/group", func(r chi.Router) {
		r.Post("/", group.Create)
		r.Get("/", group.GetList)
		r.Get("/{id}", group.GetDetail)
		r.Patch("/{id}", group.Update)
		r.Delete("/{id}", group.Delete)
	})

	return r
}
