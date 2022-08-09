package wajibPajak

import (
	"net/http"

	core "github.com/bapenda-kota-malang/apin-backend/internal/core"
	hj "github.com/bapenda-kota-malang/apin-backend/internal/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/internal/core/types"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting wajibpajak list of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := core.GetParam("id", r)
	data := t.II{
		"message": "You are visiting wajibpajak detail for id " + id + " of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are updateing wajibpajak of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
