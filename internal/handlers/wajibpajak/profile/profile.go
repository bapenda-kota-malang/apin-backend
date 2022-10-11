package profile

import (
	"errors"
	"net/http"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"

	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	su "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajak"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

// func GetList(w http.ResponseWriter, r *http.Request) {
// 	var input m.CreateDto
// 	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
// 		return
// 	}
// 	result, err := s.GetList(input)
// 	hh.DataResponse(w, result, err)
// }

// Return data detail Objek Pajak
func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.RegistrasiWajibPajak
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Create(data)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

func Checker(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	var result interface{}
	var err error
	switch id {
	case 1:
		var data m.CheckerPOneDto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		result, err = s.CheckerPOne(data)
	case 2:
		var data m.CheckerPTwoDto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		result, err = s.CheckerPTwo(data)
	case 3:
		var data mu.CheckerPThreeDto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		if *data.Password != *data.RePassword {
			err = errors.New("kombinasi password tidak sama")
			break
		}
		data.Password = nil
		data.RePassword = nil
		result, err = su.CheckerPThree(data)
	case 4:
		var data m.CheckerPFourDto
		if !hh.ValidateStructByIOR(w, r.Body, &data) {
			return
		}
		result, err = s.CheckerPFour(data)
	default:
		result = nil
		err = errors.New("unknown id page check")
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	hh.DataResponse(w, result, err)
}
