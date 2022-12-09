package profile

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func Get(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	hj.WriteJSON(w, http.StatusOK, authInfo, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your profile data in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}
