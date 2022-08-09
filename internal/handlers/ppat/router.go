package wajibpajak

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibPajak/home"
)

func SetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(er.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(er.MethodNotAllowedResponse)

	// TODO: Use group for routing

	router.HandlerFunc(http.MethodGet, "/", home.Index)

	router.HandlerFunc(http.MethodGet, "/auth/login", auth.Login)

	router.HandlerFunc(http.MethodGet, "/account/register", account.Register)
	router.HandlerFunc(http.MethodGet, "/account/resend-confirmation", account.ResendConfirmation)
	router.HandlerFunc(http.MethodGet, "/account/confirm", account.Confirm)
	router.HandlerFunc(http.MethodGet, "/account/reset-password", account.ResetPassword)
	router.HandlerFunc(http.MethodGet, "/account/change-password", account.ChangePassword)

	return router
}
