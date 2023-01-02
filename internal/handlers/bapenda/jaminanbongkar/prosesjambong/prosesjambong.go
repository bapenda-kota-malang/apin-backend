package prosesjambong

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar/prosesjambong"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/jaminanbongkar/prosesjambong"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.Create(data, uint(authInfo.User_Id), nil)
	hh.DataResponse(w, result, err)
}
