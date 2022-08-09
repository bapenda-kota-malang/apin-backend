package auth

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/pkg/core"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/core/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are loging in app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
