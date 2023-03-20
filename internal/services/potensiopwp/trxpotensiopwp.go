package potensiopwp

import (
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"

	sbapl "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/bapl"
	sdpotensiop "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailpotensiop"
	spnarahubung "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensinarahubung"
	sppemilikiwp "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensipemilikwp"
	// mvetax "github.com/bapenda-kota-malang/apin-backend/internal/models/vendoretax"
)

func CreateTrx(input m.Input, userId int) (any, error) {
	var dataPotensiOp m.PotensiOp
	var tarifPajak *mtp.TarifPajak

	potensiOpDto := input.GetPotensiOp()

	// get tarif pajak data for calculate tax
	yearOrder := "order desc"
	omsetOpt := "lte"
	rspTp, err := stp.GetList(mtp.FilterDto{
		Rekening_Id:   &potensiOpDto.Rekening_Id,
		Tahun_Opt:     &yearOrder,
		OmsetAwal:     potensiOpDto.OmsetOp,
		OmsetAwal_Opt: &omsetOpt,
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mencari tarif pajak: "+err.Error(), dataPotensiOp)
	}
	if len(rspTp.(rp.OK).Data.([]mtp.TarifPajak)) != 0 {
		tarifPajak = &rspTp.(rp.OK).Data.([]mtp.TarifPajak)[0]
	}

	// calculate tax
	if potensiOpDto, err = input.CalculateTax(tarifPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menghitung potensi pajak: "+err.Error(), dataPotensiOp)
	}

	// Transaction save to db
	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		// save potensi op
		respPotensiOp, err := Create(potensiOpDto, uint(userId), tx)
		if err != nil {
			return err
		}
		dataPotensiOp = respPotensiOp.(rp.OKSimple).Data.(m.PotensiOp)

		detailPotensiOpDto := input.GetDetailPotensiOp()
		respExistingDetailPotensiOp, err := sdpotensiop.GetExisting(
			detailPotensiOpDto.Nama,
			detailPotensiOpDto.Alamat,
			detailPotensiOpDto.RtRw,
			detailPotensiOpDto.Kecamatan_Id,
			detailPotensiOpDto.Kelurahan_Id,
			potensiOpDto.Rekening_Id,
			tx)
		if err != nil && err.Error() != "record not found" {
			return err
		}
		if respExistingDetailPotensiOp.Potensiop_Id != uuid.Nil {
			input.SetPotensiOpId(respExistingDetailPotensiOp.Potensiop_Id)
			detailPotensiOpDto = input.GetDetailPotensiOp()
		}

		// replace potensi op id
		input.ReplacePotensiOpId(dataPotensiOp.Id)
		detailPotensiOpDto = input.GetDetailPotensiOp()

		// save detail potensi op
		_, err = sdpotensiop.Create(detailPotensiOpDto, dataPotensiOp.Rekening_Id, tx)
		if err != nil {
			return err
		}

		// save pemilik wps
		_, err = sppemilikiwp.Create(input.GetPotensiPemilikWps(), tx)
		if err != nil {
			return err
		}

		_, err = spnarahubung.Create(input.GetPotensiNarahubungs(), tx)
		if err != nil {
			return err
		}

		if err := input.SaveDetailPotensiPajak(tx); err != nil {
			return err
		}

		// _, err = sdetailobjek.Create(input.DetailPajakDtos, tx)
		// if err != nil {
		// 	return err
		// }

		_, err = sbapl.Create(input.GetBapl(), uint(userId), tx)
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

func UpdateTrx(id uuid.UUID, input m.UpdateDto, userId uint) (any, error) {
	var dataPotensiOp *m.PotensiOp
	effected := 0

	// transaction update db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error

		respPotensi, err := Update(id, input.PotensiOp, userId, tx)
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

		// if len(input.DetailPajakDtos) > 0 {
		// 	_, err = sdetailobjek.Update(id, input.DetailPajakDtos, tx)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		_, err = sbapl.Update(id, input.Bapl, tx)
		if err != nil {
			return err
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
