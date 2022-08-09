package main

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/core"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda"
)

func main() {
	core.Run("apin", bapenda.SetRoutes())
}
