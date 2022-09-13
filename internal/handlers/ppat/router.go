package ppat

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
)

func SetRoutes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(er.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(er.MethodNotAllowedResponse)

	// TODO: Use group for routing

	router.HandlerFunc(http.MethodGet, "/auth/login", auth.Login)

	return router
}
