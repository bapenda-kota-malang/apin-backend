package home

import (
	"net/http"

	core "github.com/bapenda-kota-malang/apin-backend/pkg/core"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/core/types"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting list user of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
