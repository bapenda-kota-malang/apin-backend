package main

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/bapenda"
)

func main() {
	apicore.Run(bapenda.SetRoutes())
}
