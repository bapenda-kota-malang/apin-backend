package sts

import (
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sts"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func generateNomorUrut() int {
	var tmp int
	var tmpSts m.Sts
	nomor := a.DB.Last(&tmpSts)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = tmpSts.NomorUrut
		tmp++
	}
	return tmp
}

func generateNomorOutput(urut int) string {
	tmpUrut := strconv.Itoa(urut)
	if len(tmpUrut) == 1 {
		tmpUrut = "00000" + tmpUrut
	} else if len(tmpUrut) == 2 {
		tmpUrut = "0000" + tmpUrut
	} else if len(tmpUrut) == 3 {
		tmpUrut = "000" + tmpUrut
	} else if len(tmpUrut) == 4 {
		tmpUrut = "00" + tmpUrut
	} else if len(tmpUrut) == 5 {
		tmpUrut = "0" + tmpUrut
	}
	output := tmpUrut
	return output
}
