package profile

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pegawai"
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
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetDetail(authInfo.Ref_Id)
	hh.DataResponse(w, result, err)
}
