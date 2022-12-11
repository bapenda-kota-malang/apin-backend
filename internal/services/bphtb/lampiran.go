package sptpd

import (
	"sync"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func CreateLampiran(input m.CreateLampiranDto, userId uint, tx *gorm.DB) (data m.Lampiran, err error) {
	// TODO: lampiran section
	id, err := sh.GetUuidv4()
	if err != nil {
		return
	}

	var errChan = make(chan error, 16)
	var wg sync.WaitGroup

	// scan sppt
	wg.Add(1)
	fileName, path, extFile, _, err := sh.FilePreProcess(input.LampiranSppt, "lampiranSppt", userId, id)
	if err != nil {
		return
	}
	go sh.BulkSaveFile(&wg, input.LampiranSppt, fileName, path, extFile, errChan)
	input.LampiranSppt = fileName

	// fc identitas
	wg.Add(1)
	fileName, path, extFile, _, err = sh.FilePreProcess(input.LampiranFotocopiIdentitas, "lampiranFotocopiIdentitas", userId, id)
	if err != nil {
		return
	}
	go sh.BulkSaveFile(&wg, input.LampiranFotocopiIdentitas, fileName, path, extFile, errChan)
	input.LampiranFotocopiIdentitas = fileName

	// save lampiran surat pernyataan
	wg.Add(1)
	fileName, path, extFile, _, err = sh.FilePreProcess(input.LampiranFotocopySuratPernyataan, "lampiranFcSuratPernyataan", userId, id)
	if err != nil {
		return
	}
	go sh.BulkSaveFile(&wg, input.LampiranFotocopySuratPernyataan, fileName, path, extFile, errChan)
	input.LampiranFotocopySuratPernyataan = fileName

	// fc sertifikat / keterangan kepemilikan
	wg.Add(1)
	fileName, path, extFile, _, err = sh.FilePreProcess(input.LampiranSertifikatKepemilikanTanah, "lampiranSertifikatKepemilikanTanah", userId, id)
	if err != nil {
		return
	}
	go sh.BulkSaveFile(&wg, input.LampiranSertifikatKepemilikanTanah, fileName, path, extFile, errChan)
	input.LampiranSertifikatKepemilikanTanah = fileName

	// surat kuasa WP
	// TODO: ASK ABOUT THIS, SS WEBSITE NO BROWSE FILE BUTTON SO IS THIS FILE UPLOAD(?)

	// FC IDENTITAS KUASA WP
	if input.LampiranFotocopyIdentitasKwp != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranFotocopyIdentitasKwp, "lampiranFotocopyIdentitasKwp", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranFotocopyIdentitasKwp, fileName, path, extFile, errChan)
		input.LampiranFotocopyIdentitasKwp = &fileName
	}

	// fc kartu npwpw
	if input.LampiranFotocopyKartuNpwp != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranFotocopyKartuNpwp, "lampiranFotocopyKartuNpwp", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranFotocopyKartuNpwp, fileName, path, extFile, errChan)
		input.LampiranFotocopyKartuNpwp = &fileName
	}

	// akta jb
	if input.LampiranFotocopyAktaJb != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranFotocopyAktaJb, "lampiranFotocopyAktaJb", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranFotocopyAktaJb, fileName, path, extFile, errChan)
		input.LampiranFotocopyAktaJb = &fileName
	}

	// fc keterangan waris
	if input.LampiranFotocopyKeteranganWaris != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranFotocopyKeteranganWaris, "lampiranFotocopyKeteranganWaris", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranFotocopyKeteranganWaris, fileName, path, extFile, errChan)
		input.LampiranFotocopyKeteranganWaris = &fileName
	}

	// fc form spop/lspop
	if input.LampiranSpoplspop != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranSpoplspop, "lampiranSpoplspop", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranSpoplspop, fileName, path, extFile, errChan)
		input.LampiranSpoplspop = &fileName
	}

	// identitas lainnya
	if input.LampiranIdentitasLainya != nil {
		wg.Add(1)
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.LampiranIdentitasLainya, "ampiranIdentitasLainya", userId, id)
		if err != nil {
			return
		}
		go sh.BulkSaveFile(&wg, *input.LampiranIdentitasLainya, fileName, path, extFile, errChan)
		input.LampiranIdentitasLainya = &fileName
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
	if result := tx.Create(&data); result.Error != nil {
		err = result.Error
	}
	return
}
