package account

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/core"
	hj "github.com/bapenda-kota-malang/apin-backend/internal/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/internal/core/types"
)

func Register(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ResendConfirmation(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
