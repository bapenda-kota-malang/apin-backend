package bphtb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Lampiran struct {
	Id                                 uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BphtbSptpd_Id                      uuid.UUID       `json:"bphtbSptpd_id" gorm:"type:uuid"`
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
	LampiranFotocopySuratPernyataan    string          `json:"lampiranFotocopySuratPernyataan" gorm:"type:varchar(100)"`
	LampiranSpoplspop                  *string         `json:"lampiranSpoplspop" gorm:"type:varchar(110)"`
	LampiranLainnya                    *string         `json:"lampiranLainnya" gorm:"type:varchar(100)"`
	PenelitianDataObjek                *string         `json:"penelitianDataObjek" gorm:"type:varchar(1)"`
	PenelitianNilaiBphtb               *string         `json:"penelitianNilaiBphtb" gorm:"type:varchar(1)"`
	PenelitianDokumen                  *string         `json:"penelitianDokumen" gorm:"type:varchar(1)"`
	PegawaiId                          *int64          `json:"pegawaiId" gorm:"type:bigint"`
	gormhelper.DateModel
}

type CreateLampiranDto struct {
	BphtbSptpd_Id                      uuid.UUID       `json:"-"`
	NoSspd                             *string         `json:"noSspd"`
	TanggalNoSspd                      *datatypes.Date `json:"tanggalNoSspd"`
	NoFormulir                         *string         `json:"noFormulir"`
	TanggalNoFormulir                  *datatypes.Date `json:"tanggalNoFormulir"`
	LampiranSspd                       *string         `json:"lampiranSspd"`
	LampiranSppt                       string          `json:"lampiranSppt" validate:"required"`
	LampiranFotocopiIdentitas          string          `json:"lampiranFotocopiIdentitas" validate:"required"`
	LampiranSuratKuasaWp               *string         `json:"lampiranSuratKuasaWp"`
	LampiranNamaKuasaWp                *string         `json:"lampiranNamaKuasaWp"`
	LampiranAlamatKuasaWp              *string         `json:"lampiranAlamatKuasaWp"`
	LampiranFotocopyIdentitasKwp       *string         `json:"lampiranFotocopyIdentitasKwp"`
	LampiranFotocopyKartuNpwp          *string         `json:"lampiranFotocopyKartuNpwp"`
	LampiranFotocopyAktaJb             *string         `json:"lampiranFotocopyAktaJb"`
	LampiranSertifikatKepemilikanTanah string          `json:"lampiranSertifikatKepemilikanTanah" validate:"required"`
	LampiranFotocopyKeteranganWaris    *string         `json:"lampiranFotocopyKeteranganWaris"`
	LampiranFotocopySuratPernyataan    string          `json:"lampiranFotocopySuratPernyataan" validate:"required"`
	LampiranSpoplspop                  *string         `json:"lampiranSpoplspop"`
	LampiranLainnya                    *string         `json:"lampiranLainnya"`
	LampiranIdentitasLainya            *string         `json:"lampiranIdentitasLainya"`
	LampiranIdentitasLainyaVal         *string         `json:"lampiranIdentitasLainyaVal"`
	PenelitianDataObjek                *string         `json:"penelitianDataObjek"`
	PenelitianNilaiBphtb               *string         `json:"penelitianNilaiBphtb"`
	PenelitianDokumen                  *string         `json:"penelitianDokumen"`
	PegawaiId                          *int64          `json:"pegawaiId"`
}
