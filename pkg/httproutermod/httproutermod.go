// HttpRouterModified, embeded (baca: extend) httprouter
// Tujuan utama mempersingkat handling http request dengan tetap memanfaatkan
// handlerfunc untuk keperluan penggunaan middleware
// lihat variadic args pada masing2 fungsi MOD
package httproutermod

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Adapter func(http.Handler) http.Handler

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

func (r *RouterMod) GETMOD(path string, handler http.HandlerFunc, adapters ...Adapter) {
	if len(adapters) == 1 {
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
