package wajibpajak

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/profile"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/auth/login",
		"/auth/logout",
		"/register",
		"/account/register",
		"/account/reset-password",
	}
	auth.Position = 3

	// Start routing
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(auth.GuardMW)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)
	r.Route("/register", func(r chi.Router) {
		r.Post("/", profile.Create)
		r.Post("/checker/{id}", profile.Checker)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		// r.Post("/register", account.Create) // replaced withr register
		r.Get("/check", account.Check)
		r.Get("/confirm-by-email", account.ConfirmByEmail)
		r.Get("/resend-confirmation", account.ResendConfirmation)
		r.Patch("/change-pass", account.ChangePassword)
		r.Patch("/reset-pass", account.ResetPassword)
	})

	r.Route("/profile", func(r chi.Router) {
		// 	r.Get("/", p.Login)   // EXECUTE
		r.Patch("/{id}", profile.Update)
	})

	r.Route("/provinsi", func(r chi.Router) {
		r.Post("/", provinsi.Create)
		r.Get("/", provinsi.GetList)
		r.Get("/{id}", provinsi.GetDetail)
		r.Patch("/{id}", provinsi.Update)
		r.Delete("/{id}", provinsi.Delete)
	})

	r.Route("/daerah", func(r chi.Router) {
		r.Post("/", daerah.Create)
		r.Get("/", daerah.GetList)
		r.Get("/{id}", daerah.GetDetail)
		r.Patch("/{id}", daerah.Update)
		r.Delete("/{id}", daerah.Delete)
	})

	r.Route("/kecamatan", func(r chi.Router) {
		r.Post("/", kecamatan.Create)
		r.Get("/", kecamatan.GetList)
		r.Get("/{id}", kecamatan.GetDetail)
		r.Patch("/{id}", kecamatan.Update)
		r.Delete("/{id}", kecamatan.Delete)
	})

	r.Route("/kelurahan", func(r chi.Router) {
		r.Post("/", kelurahan.Create)
		r.Get("/", kelurahan.GetList)
		r.Get("/{id}", kelurahan.GetDetail)
		r.Patch("/{id}", kelurahan.Update)
		r.Delete("/{id}", kelurahan.Delete)
	})

	return r
}
