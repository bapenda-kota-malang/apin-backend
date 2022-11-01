package spt

import (
	"fmt"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
)

// Service create business flow for sptd via wajib pajak for lapor spt
//
// function flow is:
//
// create for esptd, replace id, create for data details based on data type, assign data details to data espt for respond
func CreateDetail(input m.CreateInput, user_Id uint) (interface{}, error) {
	var data m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// process
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", fmt.Sprintf("gagal mengambil menyimpan data: %s", err), data)
	}
	return rp.OKSimple{Data: data}, nil
}
