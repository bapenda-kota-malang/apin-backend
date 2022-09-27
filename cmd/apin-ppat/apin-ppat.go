package main

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/ppat"
)

func main() {
	apicore.Run(ppat.SetRoutes(), "apin/ppat")
}
