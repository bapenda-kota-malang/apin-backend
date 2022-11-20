package pengurangan

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

const source = "pengurangan"

func Create(input m.PenguranganCreateDto, user_Id uint64) (any, error) {
	var dataPengurangan m.Pengurangan
	var resp t.II
	var errChan = make(chan error)
	var baseDocsName = "pengurangan"

	fileName, path, extFile, err := FilePreProcess(*input.FotoKtp, uint(user_Id), baseDocsName+"FotoKtp")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.SaveFile(*input.FotoKtp, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// data pengurangan
	if err := sc.Copy(&dataPengurangan, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload pengurangan", dataPengurangan)
	}

	dataPengurangan.FotoKtp = &fileName
	err = a.DB.Transaction(func(tx *gorm.DB) error {
		// static value
		dataPengurangan.Pemohon_User_Id = &user_Id
		dataPengurangan.TanggalPengajuan = th.TimeNow()

		// add data file
		slcLaporanKeuangan, err := sh.GetAllTypeFile(*input.LaporanKeuangan, baseDocsName+"LaporanKeuangan", uint(user_Id))
		if err != nil {
			return err
		}
		dataPengurangan.LaporanKeuangan = &slcLaporanKeuangan
		slcLaporanPengeluaran, err := sh.GetAllTypeFile(*input.LaporanPengeluaran, baseDocsName+"LaporanPengeluaran", uint(user_Id))
		if err != nil {
			return err
		}
		dataPengurangan.LaporanPengeluaran = &slcLaporanPengeluaran
		if input.DokumenLainnya != nil {
			slcDokumenLainnya, err := sh.GetAllTypeFile(*input.DokumenLainnya, baseDocsName+"DokumenLainnya", uint(user_Id))
			if err != nil {
				return err
			}
			dataPengurangan.DokumenLainnya = &slcDokumenLainnya
		}
		slcSuratPermohonan, err := sh.GetPdfOrImageFile(*input.SuratPermohonan, baseDocsName+"SuratPermohonan", uint(user_Id))
		if err != nil {
			return err
		}
		dataPengurangan.SuratPermohonan = &slcSuratPermohonan
		// create data pengurangan
		err = tx.Create(&dataPengurangan).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataPengurangan)
	}

	resp = t.II{
		"pengurangan": dataPengurangan,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}
