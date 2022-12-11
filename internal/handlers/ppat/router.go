package ppat

import (
	"net/http"

	bphtbsptpd "github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bphtb"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/bphtbjenislaporan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/auth/login",
		"/auth/logout",
	}
	auth.Position = 2

	// Start routing
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(auth.GuardMW)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", auth.Login)
	})

	r.Route("/account", func(r chi.Router) {
		// r.Post("/register", account.Create) // replaced withr register
		r.Get("/check", account.Check)
		r.Patch("/reset-password", account.ResetPassword)
		r.Patch("/change-password", account.ChangePassword)
	})

	r.Route("/bphtbsptpd", func(r chi.Router) {
		bphtbsptpdCrud := bphtbsptpd.Crud{}
		r.Post("/", bphtbsptpdCrud.CreateMw(http.HandlerFunc(bphtbsptpdCrud.Create), "ppat"))
		r.Get("/", bphtbsptpdCrud.GetList)
		r.Get("/{id}", bphtbsptpdCrud.GetDetail)
		r.Patch("/{id}/verify", bphtbsptpdCrud.VerifyPpat)
		r.Get("/jenislaporan", bphtbjenislaporan.Crud{}.GetList)
	})

	return r
}
