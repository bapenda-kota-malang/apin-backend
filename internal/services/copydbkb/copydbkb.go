package copydbkb

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/copydbkb"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"gorm.io/gorm"
)

const source = "CopyDbkb"

// this function will call SP Penetapan Massal
func SpCopyDBKB(input m.CopyDBKBDto, tx *gorm.DB) (*m.SpCopyDbkb, error) {
	if tx == nil {
		tx = a.DB
	}
	var (
		rsltSp *m.SpCopyDbkb
	)

	// execute SpCopyDbkb
	res := tx.Raw("SELECT * FROM CopyDbkb(?)",
		input.Tahun).Scan(&rsltSp)
	if res.Error != nil {
		return nil, res.Error
	}

	return rsltSp, nil
}

func CopyDBKB(input m.CopyDBKBDto) (any, error) {
	var res *m.SpCopyDbkb
	resSp, err := SpCopyDBKB(input, a.DB)
	if err != nil {
		return sh.SetError("request", "SpCopyDbkb", source, "failed", "data tidak ditemukan", res)
	}

	return rp.OKSimple{Data: resSp}, nil
}
