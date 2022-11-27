package pengurangan

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	ms "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

const source = "keberatan"

func Create(input m.KeberatanCreateDto, user_Id uint64) (any, error) {
	var dataKeberatan m.Keberatan
	var resp t.II
	var errChan = make(chan error)
	var baseDocsName = "keberatan"

	fileName, path, extFile, err := ms.FilePreProcess(*input.FotoKtp, uint(user_Id), baseDocsName+"FotoKtp")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(*input.FotoKtp, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// data keberatan
	if err := sc.Copy(&dataKeberatan, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload keberatan", dataKeberatan)
	}

	dataKeberatan.FotoKtp = &fileName
	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// static value
		dataKeberatan.Pemohon_User_Id = &user_Id
		dataKeberatan.TanggalPengajuan = th.TimeNow()

		// add data file
		// slcLaporanKeuangan, err := sh.GetAllTypeFile(*input.LaporanKeuangan, baseDocsName+"LaporanKeuangan", uint(user_Id))
		// if err != nil {
		// 	return err
		// }
		// dataKeberatan.LaporanKeuangan = &slcLaporanKeuangan
		// slcLaporanPengeluaran, err := sh.GetAllTypeFile(*input.LaporanPengeluaran, baseDocsName+"LaporanPengeluaran", uint(user_Id))
		// if err != nil {
		// 	return err
		// }
		// dataKeberatan.LaporanPengeluaran = &slcLaporanPengeluaran
		if input.DokumenLainnya != nil {
			slcDokumenLainnya, err := sh.GetAllTypeFile(*input.DokumenLainnya, baseDocsName+"DokumenLainnya", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.DokumenLainnya = &slcDokumenLainnya
		}
		if input.SuratPermohonan != nil {
			slcSuratPermohonan, err := sh.GetPdfOrImageFile(*input.SuratPermohonan, baseDocsName+"SuratPermohonan", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.SuratPermohonan = &slcSuratPermohonan
		}
		if input.SuratPernyataan != nil {
			slcSuratPernyataan, err := sh.GetPdfOrImageFile(*input.SuratPernyataan, baseDocsName+"SuratPernyataan", uint(user_Id))
			if err != nil {
				return err
			}
			dataKeberatan.SuratPernyataan = &slcSuratPernyataan
		}
		// create data keberatan
		err = tx.Create(&dataKeberatan).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataKeberatan)
	}

	resp = t.II{
		"keberatan": dataKeberatan,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
