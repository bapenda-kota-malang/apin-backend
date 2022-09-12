package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/group"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/pendaftaran"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
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

	r.Route("/registration", func(r chi.Router) {
		r.Get("/", pendaftaran.GetList)
		r.Get("/{id}", pendaftaran.GetDetail)
		r.Post("/operator", pendaftaran.RegisterByOperator)
	})

	r.Route("/rekening", func(r chi.Router) {
		r.Get("/", rekening.GetList)
	})

	return r
}
