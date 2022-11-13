package wajibpajak

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/configuration/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/daerah"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kecamatan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/kelurahan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda/provinsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/account"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/auth"
	er "github.com/bapenda-kota-malang/apin-backend/internal/handlers/main/errors"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/hargadasarair"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/home"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/profile"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/regnpwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak/tarifpajak"
	rh "github.com/bapenda-kota-malang/apin-backend/pkg/routerhelper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetRoutes() http.Handler {
	// Config
	auth.SkipAuhPaths = []string{
		"/auth/login",
		"/auth/logout",
		"/register",
		"/account/register",
		"/account/reset-password",
		"/account/change-password",
		"/provinsi",
		"/daerah",
		"/kecamatan",
		"/kelurahan",
		"/rekening",
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
		r.Get("/", profile.GetDetail) // EXECUTE
		r.Patch("/{id}", profile.Update)
	})

	r.Route("/provinsi", func(r chi.Router) {
		r.Get("/", provinsi.GetList)
		r.Get("/{id}", provinsi.GetDetail)
	})

	r.Route("/daerah", func(r chi.Router) {
		r.Get("/", daerah.GetList)
		r.Get("/{id}", daerah.GetDetail)
	})

	r.Route("/kecamatan", func(r chi.Router) {
		r.Get("/", kecamatan.GetList)
		r.Get("/{id}", kecamatan.GetDetail)
	})

	r.Route("/kelurahan", func(r chi.Router) {
		r.Get("/", kelurahan.GetList)
		r.Get("/{id}", kelurahan.GetDetail)
	})

	r.Route("/regnpwpd", func(r chi.Router) {
		r.Post("/", regnpwpd.Create)
		r.Get("/", regnpwpd.GetListForWp)
		r.Get("/{id}", regnpwpd.GetDetailForWp)
		r.Patch("/{id}", regnpwpd.UpdateForWp)
		r.Delete("/{id}", regnpwpd.DeleteForWp)
	})

	r.Route("/espt", func(r chi.Router) {
		r.Post("/", espt.Create)
		r.Get("/", espt.GetList)
		r.Get("/{id}", espt.GetDetail)
		r.Patch("/{id}", espt.Update)
		r.Delete("/{id}", espt.Delete)
	})

	rh.RegCrud(r, "/rekening", rekening.Crud{})

	r.Route("/npwpd", func(r chi.Router) {
		r.Get("/", npwpd.GetList)
		r.Get("/{id}", npwpd.GetDetail)
		r.Patch("/{id}", npwpd.Update)
		r.Patch("/{id}/{category}", npwpd.UpdatePhoto)
		r.Delete("/{id}/{category}/{filename}", npwpd.DeletePhoto)
	})

	r.Route("/jenisppj", func(r chi.Router) {
		r.Get("/", jenisppj.GetList)
		r.Get("/{id}", jenisppj.GetDetail)
	})

	r.Route("/tarifpajak", func(r chi.Router) {
		r.Get("/", tarifpajak.GetList)
	})

	r.Route("/hargadasarair", func(r chi.Router) {
		r.Get("/", hargadasarair.GetList)
		r.Get("/peruntukan", hargadasarair.GetPeruntukan)
	})

	r.Route("/sptpd", func(r chi.Router) {
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
	})

	r.Route("/skpd", func(r chi.Router) {
		r.Get("/", spt.GetList)
		r.Get("/{id}", spt.GetDetail)
	})

	r.Route("/pengajuan", func(r chi.Router) {
		r.Post("/{category}", pengurangan.Create)
	})
	return r
}
