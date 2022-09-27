package auth

import (
	"fmt"
	"net/http"
	"strconv"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	as "github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var input um.LoginDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	input.Position = 2
	result, err := as.GenToken(input)
	if err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		fmt.Println(err)
		hj.WriteJSON(w, http.StatusUnauthorized, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(1)},
			Messages: err.Error(),
		}, nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
