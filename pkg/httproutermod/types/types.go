package types

import "net/http"

// type Adapter func(http.Handler) http.Handler

type HandlerFuncAdapter func(http.HandlerFunc) http.HandlerFunc
