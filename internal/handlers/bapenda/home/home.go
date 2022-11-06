package home

import (
	"net/http"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "Selamat datang di API Integrasi Pajak Daerah Bapenda Kota Malang",
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
