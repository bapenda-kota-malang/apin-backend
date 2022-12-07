package bphtb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Lampiran struct {
	Id                                 uint            `json:"id" gorm:"primarykey"`
	NoSspd                             *string         `json:"noSspd" gorm:"type:varchar(20)"`
	TanggalNoSspd                      *datatypes.Date `json:"tanggalNoSspd"`
	NoFormulir                         *string         `json:"noFormulir" gorm:"type:varchar(20)"`
	TanggalNoFormulir                  *datatypes.Date `json:"tanggalNoFormulir"`
	LampiranSspd                       *string         `json:"lampiranSspd" gorm:"type:varchar(100)"`
	LampiranSppt                       *string         `json:"lampiranSppt" gorm:"type:varchar(100)"`
	LampiranFotocopiIdentitas          *string         `json:"lampiranFotocopiIdentitas" gorm:"type:varchar(100)"`
	LampiranSuratKuasaWp               *string         `json:"lampiranSuratKuasaWp" gorm:"type:varchar(1)"`
	LampiranNamaKuasaWp                *string         `json:"lampiranNamaKuasaWp" gorm:"type:varchar(200)"`
	LampiranAlamatKuasaWp              *string         `json:"lampiranAlamatKuasaWp" gorm:"type:varchar(200)"`
	LampiranFotocopyIdentitasKwp       *string         `json:"lampiranFotocopyIdentitasKwp" gorm:"type:varchar(100)"`
	LampiranFotocopyKartuNpwp          *string         `json:"lampiranFotocopyKartuNpwp" gorm:"type:varchar(100)"`
	LampiranIdentitasLainya            *string         `json:"lampiranIdentitasLainya" gorm:"type:varchar(100)"`
	LampiranIdentitasLainyaVal         *string         `json:"lampiranIdentitasLainyaVal" gorm:"type:varchar(100)"`
	LampiranFotocopyAktaJb             *string         `json:"lampiranFotocopyAktaJb" gorm:"type:varchar(100)"`
	LampiranSertifikatKepemilikanTanah *string         `json:"lampiranSertifikatKepemilikanTanah" gorm:"type:varchar(110)"`
	LampiranFotocopyKeteranganWaris    *string         `json:"lampiranFotocopyKeteranganWaris" gorm:"type:varchar(100)"`
	LampiranFotocopySuratPernyataan    *string         `json:"lampiranFotocopySuratPernyataan" gorm:"type:varchar(100)"`
	LampiranSpoplspop                  *string         `json:"lampiranSpoplspop" gorm:"type:varchar(110)"`
	LampiranLainnya                    *string         `json:"lampiranLainnya" gorm:"type:varchar(100)"`
	PenelitianDataObjek                *string         `json:"penelitianDataObjek" gorm:"type:varchar(1)"`
	PenelitianNilaiBphtb               *string         `json:"penelitianNilaiBphtb" gorm:"type:varchar(1)"`
	PenelitianDokumen                  *string         `json:"penelitianDokumen" gorm:"type:varchar(1)"`
	PegawaiId                          *string         `json:"pegawaiId" gorm:"type:bigint(20)"`
	gormhelper.DateModel
}
