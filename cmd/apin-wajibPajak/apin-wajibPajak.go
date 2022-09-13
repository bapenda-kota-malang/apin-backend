package main

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibpajak"
)

func main() {
	apicore.Run(wajibpajak.SetRoutes())
}
