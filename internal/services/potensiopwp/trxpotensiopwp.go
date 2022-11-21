package potensiopwp

import (
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"

	sdetailobjek "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailobjek"
	sdpotensiop "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailpotensiop"
	spnarahubung "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensinarahubung"
	sppemilikiwp "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensipemilikwp"
	// mvetax "github.com/bapenda-kota-malang/apin-backend/internal/models/vendoretax"
)

func CreateTrx(input m.CreateDto, userId uint) (any, error) {
	var dataPotensiOp m.PotensiOp

	input.PotensiOp.User_Id = userId

	// Transaction save to db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		// save potensi op
		respPotensiOp, err := Create(input.PotensiOp, tx)
		if err != nil {
			return err
		}
		dataPotensiOp = respPotensiOp.(rp.OKSimple).Data.(m.PotensiOp)

		// replace potensi op id
		input.DetailPotensiOp.Potensiop_Id = dataPotensiOp.Id
		for v := range input.PotensiPemilikWps {
			input.PotensiPemilikWps[v].Potensiop_Id = dataPotensiOp.Id
		}
		for v := range input.PotensiNarahubungs {
			input.PotensiNarahubungs[v].Potensiop_Id = dataPotensiOp.Id
		}
		for v := range input.DetailPajakDtos {
			input.DetailPajakDtos[v].Potensiop_Id = dataPotensiOp.Id
		}

		// save detail potensi op
		_, err = sdpotensiop.Create(input.DetailPotensiOp, tx)
		if err != nil {
			return err
		}

		// save pemilik wps
		_, err = sppemilikiwp.Create(input.PotensiPemilikWps, tx)
		if err != nil {
			return err
		}

		// save narahubungs if len narahubung > 0
		if len(input.PotensiNarahubungs) > 0 {
			_, err := spnarahubung.Create(input.PotensiNarahubungs, tx)
			if err != nil {
				return err
			}
		}

		_, err = sdetailobjek.Create(input.DetailPajakDtos, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data potensi: "+err.Error(), dataPotensiOp)
	}

	return rp.OKSimple{Data: dataPotensiOp}, nil
}

func UpdateTrx(id int, input m.UpdateDto, userId uint) (any, error) {
	var dataPotensiOp *m.PotensiOp
	effected := 0

	// transaction update db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error

		input.PotensiOp.User_Id = &userId
		respPotensi, err := Update(id, input.PotensiOp, tx)
		if err != nil {
			return err
		}
		dataPotensiOp = respPotensi.(rp.OK).Data.(*m.PotensiOp)
		affectedDataPotensi, _ := strconv.Atoi(respPotensi.(rp.OK).Meta["affected"])
		effected += affectedDataPotensi

		// save detail potensi op
		_, err = sdpotensiop.Update(id, input.DetailPotensiOp, tx)
		if err != nil {
			return err
		}

		// save pemilik wps
		if len(input.PotensiPemilikWps) > 0 {
			_, err = sppemilikiwp.Update(id, input.PotensiPemilikWps, tx)
			if err != nil {
				return err
			}
		}

		// save narahubungs if len narahubung > 0
		if len(input.PotensiNarahubungs) > 0 {
			_, err := spnarahubung.Update(id, input.PotensiNarahubungs, tx)
			if err != nil {
				return err
			}
		}

		if len(input.DetailPajakDtos) > 0 {
			_, err = sdetailobjek.Update(id, input.DetailPajakDtos, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data potensi: "+err.Error(), dataPotensiOp)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": "1",
		},
		Data: dataPotensiOp,
	}, nil
}
