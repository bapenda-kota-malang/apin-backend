package main

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/core"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/wajibPajak"
)

func main() {
	core.Run(wajibPajak.SetRoutes())
}
