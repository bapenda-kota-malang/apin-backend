package wajibpajak

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/home"
	// "github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/profile" // EXECUTE
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)
	// r.Post("/register", wp.Create) // EXECUTE

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		// r.Post("/register", account.Create) // replaced withr register
		r.Get("/confirm-by-email", account.ConfirmByEmail)
		r.Get("/resend-confirmation", account.ResendConfirmation)
		r.Patch("/change-pass", account.ChangePassword)
		r.Patch("/reset-pass", account.ResetPassword)
	})

	// r.Route("/profile", func(r chi.Router) { // EXECUTE
	// 	r.Get("/", p.Login)   // EXECUTE
	// 	r.Patch("/", p.Login) // EXECUTE
	// })

	return r
}
