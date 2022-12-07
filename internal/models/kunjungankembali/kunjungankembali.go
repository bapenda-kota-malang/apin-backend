package kunjungankembali

import (
	"time"

	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type KunjunganKembali struct {
	opp.NopDetail
	NoBangunan              int        `json:"noBangunan"`
	TanggalKunjunganKembali *time.Time `json:"tanggalKunjunganKembali"`
	Status                  int        `json:"number"`
	gh.DateModel
}
