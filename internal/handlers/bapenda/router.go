package bapenda

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
)

func SetRoutes() http.Handler {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.NotFound(er.NotFoundResponse)
	r.MethodNotAllowed(er.MethodNotAllowedResponse)

	r.Get("/", home.Index)

	r.Post("/auth/login", auth.Login)

	r.Patch("/account/reset-password", account.ResetPassword)
	r.Patch("/account/change-password", account.ChangePassword)

	r.Post("/user", user.Create)
	r.Get("/user", user.GetList)
	r.Get("/user/{id}", user.GetDetail)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sup"))
	})

	return r
}
