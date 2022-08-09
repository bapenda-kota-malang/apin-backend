package home

import (
	"net/http"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting list user of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
