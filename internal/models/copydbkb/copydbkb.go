package copydbkb

import (
// gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type CopyDBKBDto struct {
	Tahun string `json:"tahun" validate:"requred;min=1"`
}

type SpCopyDbkb struct {
	CopyDbkb *float64 `gorm:"column:copydbkb"`
}
