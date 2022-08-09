package main

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibPajak"
)

func main() {
	apicore.Run(wajibPajak.SetRoutes())
}
