package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	js, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(500)
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
