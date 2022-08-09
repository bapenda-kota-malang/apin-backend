package userservice

import (
	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func Create(data um.UserCreate) t.II {
	return t.II{"test": "ok"}
}
