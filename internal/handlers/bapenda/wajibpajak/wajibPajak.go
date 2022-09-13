package wajibPajak

import (
	"fmt"
	"net/http"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting wajibpajak list of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	// id := .GetParam("id", r)
	data := t.II{
		"message": "You are visiting wajibpajak detail for id of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are updateing wajibpajak of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
