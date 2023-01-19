package account

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
)

func Register(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ResendConfirmation(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ConfirmByEmail(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are getting your account data in app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Check(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	data := t.II{
		"data": authInfo,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var input m.ChangePassDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	ctx := r.Context()
	authInfo, ok := ctx.Value("authInfo").(*auth.AuthInfo)
	if !ok {
		hj.WriteJSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	result, err := s.ChangePass(authInfo.User_Id, input)
	hh.DataResponse(w, result, err)
}

func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	var input m.RequestResetPassDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.RequestResetPass(input)
	hh.DataResponse(w, result, err)
}

func CheckResetPassword(w http.ResponseWriter, r *http.Request) {
	var input m.CheckResetPassDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.CheckResetPass(input)
	hh.DataResponse(w, result, err)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var input1 m.CheckResetPassDto
	if hh.ValidateStructByURL(w, *r.URL, &input1) == false {
		return
	}
	var input2 m.ResetPassDto
	if hh.ValidateStructByIOR(w, r.Body, &input2) == false {
		return
	}

	result, err := s.ResetPass(input1, input2)
	hh.DataResponse(w, result, err)
}
