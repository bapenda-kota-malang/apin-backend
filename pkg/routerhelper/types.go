package routerhelper

import "net/http"

type CrudBase interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	GetDetail(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
