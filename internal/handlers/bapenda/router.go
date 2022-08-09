package bapenda

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
)

func SetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(er.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(er.MethodNotAllowedResponse)

	// TODO: Use group for routing

	router.HandlerFunc(http.MethodGet, "/", home.Index)

	router.HandlerFunc(http.MethodPost, "/auth/login", auth.Login)

	router.HandlerFunc(http.MethodPatch, "/account/reset-password", account.ResetPassword)
	router.HandlerFunc(http.MethodPatch, "/account/change-password", account.ChangePassword)

	router.HandlerFunc(http.MethodPost, "/user/", user.Create)
	router.HandlerFunc(http.MethodGet, "/user/", user.GetList)
	router.HandlerFunc(http.MethodPatch, "/user/:id", user.GetDetail)
	return router
}
