package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PstPenguranganDendaAdm struct {
	Id                    uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64  `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string  `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string  `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string  `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string  `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string  `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string  `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string  `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string  `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string  `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string  `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string  `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string  `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	JenisPengurangan      *string  `json:"jenisPengurangan" gorm:"type:varchar(1)"`
	TahunPengurangan      *string  `json:"tahunPengurangan" gorm:"type:varchar(4)"`
	JenisSK               *string  `json:"jenisSK" gorm:"type:varchar(1)"`
	NoSK                  *string  `json:"noSK" gorm:"type:varchar(30)"`
	StatusSK              *string  `json:"statusSK" gorm:"type:varchar(1)"`
	PersentasePengurangan *float64 `json:"persentasePengurangan" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}
