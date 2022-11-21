package pengurangan

import (
	"fmt"
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan"
	sk "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan/keberatan"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	fmt.Println("category: ", category)
	switch category {
	case "pengurangan":
		var input m.PenguranganCreateDto
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.Create(input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "keberatan":
		var input m.KeberatanCreateDto
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := sk.Create(input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	default:
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "category tidak diketahui"}, nil)
		return
	}

}
