package npwpd

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	user_id := uint64(authInfo.User_Id)
	input.User_Id = &user_id

	result, err := s.GetListForWp(input)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.GetDetailByUser(id, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Update(id, input, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	category := chi.URLParam(r, "category")
	switch category {
	case "fotoktp":
		var input m.PhotoUpdate
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.AddPhotoKtp(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "lainlain":
		var input m.PhotoUpdate
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.AddPhotoLainLain(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "fotoobjek":
		var input m.PhotoUpdate
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.AddPhotoObject(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "suratizinusaha":
		var input m.PhotoUpdate
		if hh.ValidateStructByIOR(w, r.Body, &input) == false {
			return
		}

		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.AddPhotoSuratIzin(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	default:
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "category tidak diketahui"}, nil)
		return
	}
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	category := chi.URLParam(r, "category")
	input := chi.URLParam(r, "filename")
	switch category {
	case "lainlain":
		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.DeletePhotoLainLain(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "fotoobjek":
		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.DeletePhotoObject(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	case "suratizinusaha":
		authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
		result, err := s.DeletePhotoSuratIzin(id, input, uint64(authInfo.User_Id))
		hh.DataResponse(w, result, err)
	default:
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "category tidak diketahui"}, nil)
		return
	}
}
