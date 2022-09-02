package bapenda

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	amw "github.com/bapenda-kota-malang/apin-backend/pkg/authorizationmw"
	hrm "github.com/bapenda-kota-malang/apin-backend/pkg/httproutermod"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
)

func SetRoutes() *httprouter.Router {
	rmod := hrm.New()
	rmod.NotFound = http.HandlerFunc(er.NotFoundResponse)
	rmod.MethodNotAllowed = http.HandlerFunc(er.MethodNotAllowedResponse)

	// TODO: Use group for routing

	rmod.GETMOD("/", home.Index)
	rmod.GETMOD("/v1/checkauth", home.Index, amw.CheckAuth)

	rmod.POSTMOD("/v1/auth/login", auth.Login)
	rmod.POSTMOD("/v1/auth/logout", auth.Logout)

	rmod.PATCHMOD("/v1/account/reset-password", account.ResetPassword)
	rmod.PATCHMOD("/v1/account/change-password", account.ChangePassword)

	rmod.HandlerFunc(http.MethodPost, "/v1/user/", user.Create)
	rmod.HandlerFunc(http.MethodGet, "/v1/user/", user.GetList)
	rmod.HandlerFunc(http.MethodPatch, "/v1/user/:id", user.GetDetail)

	return rmod.Router
}
