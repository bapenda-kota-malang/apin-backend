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
	rmod.GETMOD("/checkauth", home.Index, amw.CheckAuth)

	rmod.POSTMOD("/auth/login", auth.Login)

	rmod.PATCHMOD("/account/reset-password", account.ResetPassword)
	rmod.PATCHMOD("/account/change-password", account.ChangePassword)

	rmod.HandlerFunc(http.MethodPost, "/user/", user.Create)
	rmod.HandlerFunc(http.MethodGet, "/user/", user.GetList)
	rmod.HandlerFunc(http.MethodPatch, "/user/:id", user.GetDetail)

	return rmod.Router
}
