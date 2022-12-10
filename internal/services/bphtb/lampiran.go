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

	// save lampiran surat pernyataan
	wg.Add(1)
	fileName, path, extFile, _, err := sh.FilePreProcess(input.LampiranFotocopySuratPernyataan, "lampiranFcSuratPernyataan", userId, id)
	if err != nil {
		return
	}
	go sh.BulkSaveFile(&wg, input.LampiranFotocopySuratPernyataan, fileName, path, extFile, errChan)
	input.LampiranFotocopySuratPernyataan = fileName

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
