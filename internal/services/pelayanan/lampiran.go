package permohonan

import (
	"strings"
	"sync"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func CreateLampiran(input m.PstLampiranCreateDTO, userId uint, tx *gorm.DB) (data m.PstLampiran, err error) {
	var (
		fileName string
		path     string
		extFile  string
	)
	// TODO: lampiran section
	id, err := sh.GetUuidv4()
	if err != nil {
		return
	}

	result := a.DB.Where("PermohonanId", input.PermohonanId).First(&data)
	data.PermohonanId = input.PermohonanId
	data.BuktiKepemilikan = input.BuktiKepemilikan

	var errChan = make(chan error, 16)
	var wg sync.WaitGroup

	if input.LampiranKTP != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranKTP, "lampiranKTP", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranKTP, fileName, path, extFile, errChan)
		data.LampiranKTP = &fileName
	}

	if input.LampiranAkte != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranAkte, "lampiranAkte", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranAkte, fileName, path, extFile, errChan)
		data.LampiranAkte = &fileName
	}

	if input.LampiranImb != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranImb, "lampiranImb", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranImb, fileName, path, extFile, errChan)
		data.LampiranImb = &fileName
	}

	if input.LampiranLainLain != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranLainLain, "lampiranLainLain", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranLainLain, fileName, path, extFile, errChan)
		data.LampiranLainLain = &fileName
	}

	if input.LampiranLikuid != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranLikuid, "lampiranLikuid", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranLikuid, fileName, path, extFile, errChan)
		input.LampiranLikuid = &fileName
	}

	if input.LampiranLaporanKeuangan != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranLaporanKeuangan, "lampiranLaporanKeuangan", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranLaporanKeuangan, fileName, path, extFile, errChan)
		input.LampiranLaporanKeuangan = &fileName
	}

	if input.LampiranSlipGaji != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSlipGaji, "lampiranSlipGaji", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSlipGaji, fileName, path, extFile, errChan)
		input.LampiranSlipGaji = &fileName
	}

	if input.LampiranPermohonan != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranPermohonan, "lampiranPermohonan", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranPermohonan, fileName, path, extFile, errChan)
		input.LampiranPermohonan = &fileName
	}

	if input.LampiranListrik != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranListrik, "lampiranListrik", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranListrik, fileName, path, extFile, errChan)
		input.LampiranListrik = &fileName
	}

	if input.LampiranSTTS != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSTTS, "lampiranSTTS", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSTTS, fileName, path, extFile, errChan)
		data.LampiranSTTS = &fileName
	}

	if input.LampiranSertifikatTanah != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSertifikatTanah, "lampiranSertifikatTanah", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSertifikatTanah, fileName, path, extFile, errChan)
		data.LampiranSertifikatTanah = &fileName
	}

	if input.LampiranSkKeberatan != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSkKeberatan, "lampiranSkKeberatan", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSkKeberatan, fileName, path, extFile, errChan)
		data.LampiranSkKeberatan = &fileName
	}

	if input.LampiranSkPengurangan != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSkPengurangan, "lampiranSkPengurangan", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSkPengurangan, fileName, path, extFile, errChan)
		data.LampiranSkPengurangan = &fileName
	}

	if input.LampiranSkPensiun != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSkPensiun, "lampiranSkPensiun", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSkPensiun, fileName, path, extFile, errChan)
		data.LampiranSkPensiun = &fileName
	}

	if input.LampiranSK != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSK, "lampiranSK", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSK, fileName, path, extFile, errChan)
		data.LampiranSK = &fileName
	}

	if input.LampiranKK != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranKK, "lampiranKK", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranKK, fileName, path, extFile, errChan)
		data.LampiranKK = &fileName
	}

	if input.LampiranSkkpPbb != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSkkpPbb, "lampiranSkkpPbb", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSkkpPbb, fileName, path, extFile, errChan)
		data.LampiranSkkpPbb = &fileName
	}

	if input.LampiranSpmkpPbb != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSpmkpPbb, "lampiranSpmkpPbb", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSpmkpPbb, fileName, path, extFile, errChan)
		data.LampiranSpmkpPbb = &fileName
	}

	if input.LampiranSppt != nil {
		tempSPPT := *input.LampiranSppt
		var tempResult []string
		for _, item := range tempSPPT {
			wg.Add(1)
			fileName, path, extFile, _, err = sh.FilePreProcess(item, "lampiranSppt", userId, id)
			if err != nil {
				return
			}
			go sh.BulkSaveFile(&wg, item, fileName, path, extFile, errChan)
			tempResult = append(tempResult, fileName)
		}
		joinResult := strings.Join(tempResult, ", ")
		data.LampiranSppt = &joinResult
	}

	if input.LampiranSpptStts != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSpptStts, "lampiranSpptStts", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSpptStts, fileName, path, extFile, errChan)
		data.LampiranSpptStts = &fileName
	}

	if input.LampiranSuratKuasa != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSuratKuasa, "lampiranSuratKuasa", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSuratKuasa, fileName, path, extFile, errChan)
		data.LampiranSuratKuasa = &fileName
	}

	if input.LampiranLetakOP != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranLetakOP, "lampiranLetakOP", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranLetakOP, fileName, path, extFile, errChan)
		data.LampiranLetakOP = &fileName
	}

	if input.LampiranFotoOP != nil {
		tempFotoOP := *input.LampiranFotoOP
		var tempResult []string
		for _, item := range tempFotoOP {
			wg.Add(1)
			fileName, path, extFile, _, err = sh.FilePreProcess(item, "lampiranFotoOP", userId, id)
			if err != nil {
				return
			}
			go sh.BulkSaveFile(&wg, item, fileName, path, extFile, errChan)
			tempResult = append(tempResult, fileName)
		}
		joinResult := strings.Join(tempResult, ", ")
		data.LampiranFotoOP = &joinResult
	}

	if input.LampiranHakMilik != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranHakMilik, "lampiranHakMilik", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranHakMilik, fileName, path, extFile, errChan)
		data.LampiranHakMilik = &fileName
	}

	if err = sc.Copy(&data, &input); err != nil {
		return
	}

	// static data
	data.Id = id

	wg.Wait()
	close(errChan)
	for v := range errChan {
		if v != nil {
			err = v
			return
		}
	}

	if result.RowsAffected > 0 {
		if result := tx.Where("PermohonanId", input.PermohonanId).Save(&data); result.Error != nil {
			err = result.Error
		}
	} else {
		if result := tx.Create(&data); result.Error != nil {
			err = result.Error
		}
	}
	return
}
