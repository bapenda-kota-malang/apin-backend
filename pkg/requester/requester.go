package requester

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetParam(param string, r *http.Request) string {
	return httprouter.ParamsFromContext(r.Context()).ByName(param)
}

func GetIntParam(param string, r *http.Request) (int64, error) {
	result, err := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName(param), 10, 64)
	if err != nil || result < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return result, nil
}
