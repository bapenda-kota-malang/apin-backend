package profile

import (
	"net/http"

	core "github.com/bapenda-kota-malang/apin-backend/pkg/core"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/core/types"
)

func Get(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your profile data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your profile data in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
