package sts

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func generateNomor() int {
	var tmp int
	var tmpSts m.Sts
	nomor := a.DB.Last(&tmpSts)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = *tmpSts.StsNumber
		tmp++
	}
	return tmp
}
