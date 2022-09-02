package auth

import (
	"net/http"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are loging in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are loging in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
