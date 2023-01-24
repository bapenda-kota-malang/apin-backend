package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegPstDataOPBaru struct {
	Id              uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId    *uint64 `json:"permohonanId" gorm:"type:integer"`
	KanwilId        *string `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId         *string `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan  *string `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan *string `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan *string `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	NamaWPBaru      *string `json:"namaWP" gorm:"type:varchar(30)"`
	LetakOPBaru     *string `json:"letakOP" gorm:"type:varchar(40)"`
	gormhelper.DateModel
}
