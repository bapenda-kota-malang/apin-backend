package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegPstPermohonanPengurangan struct {
	Id                    uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64  `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string  `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string  `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string  `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string  `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string  `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	JenisPengurangan      *string  `json:"jenisPengurangan" gorm:"type:varchar(1)"`
	AlasanPengurangan     *string  `json:"alasanPengurangan" gorm:"type:varchar(2)"`
	PersentasePengurangan *float64 `json:"persentasePengurangan" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}
