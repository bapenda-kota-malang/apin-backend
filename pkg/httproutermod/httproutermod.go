/*
HttpRouterModified, embeded (baca: extend) httprouter
Tujuan:
1. mempersingkat handling http request
2. mengijinkan penggunaan middleware, lihat variadic args pada masing2 fungsi MOD
*/
package httproutermod

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bapenda-kota-malang/apin-backend/pkg/httproutermod/types"
)

// Embed here
type RouterMod struct {
	*httprouter.Router
}

func New() *RouterMod {
	return &RouterMod{
		Router: &httprouter.Router{
			RedirectTrailingSlash:  true,
			RedirectFixedPath:      true,
			HandleMethodNotAllowed: true,
			HandleOPTIONS:          true,
		},
	}
}

/*
Berikut pattern membuat middleware

// Memanfaatkan pattern types.Adapter dari sub package types.
func CheckAuth(h http.HandlerFunc) http.HandlerFunc {
	// return sesuai handlerfunc
	return func(w http.ResponseWriter, r *http.Request) {
		// do anything here
		.....

		// serve the input `h`, put inside `if` if necessary
		h.ServeHTTP(w, r)
	}
}

*/

// func (r *RouterMod) GETMOD(path string, handler http.HandlerFunc, adapters ...types.Adapter) {
func (r *RouterMod) GETMOD(path string, handler http.HandlerFunc, adapters ...types.HandlerFuncAdapter) {
	if len(adapters) > 0 {
		// var finalHandlerFunc http.HandlerFunc = handler
		// for _, handler := range adapters {
		// 	finalHandlerFunc = handler(finalHandlerFunc)
		// }
		// r.Handler(http.MethodGet, path, finalHandlerFunc)
		r.Handler(http.MethodGet, path, adapters[0](handler))
	} else {
		r.Handler(http.MethodGet, path, handler)
	}
}

func (r *RouterMod) POSTMOD(path string, handler http.HandlerFunc) {
	r.Handler(http.MethodGet, path, handler)
}

func (r *RouterMod) PATCHMOD(path string, handler http.HandlerFunc) {
	r.Handler(http.MethodPatch, path, handler)
}

func (r *RouterMod) PUTMOD(path string, handler http.HandlerFunc) {
	r.Handler(http.MethodPut, path, handler)
}

func (r *RouterMod) DELMOD(path string, handler http.HandlerFunc) {
	r.Handler(http.MethodDelete, path, handler)
}
