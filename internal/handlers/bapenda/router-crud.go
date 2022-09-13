package bapenda

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/crudbase"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/jabatan"
)

func init() {
	crudbase.Router = *r
}

func setCrudRoutes() {
	crudbase.Register("/jabatan", jabatan.Jabatan{})
}
