package rekening

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := rekening.GetAll(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}
