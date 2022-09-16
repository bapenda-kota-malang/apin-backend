package responses

import (
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

type OK struct {
	Meta t.IS        `json:"meta"`
	Data interface{} `json:"data"`
}

type OKSimple struct {
	Data interface{} `json:"data"`
}

type OKComplex struct {
	Meta t.IS        `json:"meta"`
	Data interface{} `json:"data"`
	Ref  interface{} `json:"ref"`
}

type Err struct {
	Meta     t.IS `json:"meta"`
	Messages t.IS `json:"messages"`
}

type ErrSimple struct {
	Message string `json:"messages"`
}

type ErrCustom struct {
	Meta     interface{} `json:"meta"`
	Messages interface{} `json:"message"`
}
