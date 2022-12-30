package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SkSk struct {
	Id            uint64          `json:"id" gorm:"primaryKey"`
	PermohonanId  *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId      *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId       *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	JnsSK         *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK          *string         `json:"noSK" gorm:"type:varchar(30)"`
	TglSK         *datatypes.Date `json:"tglSK"`
	NoBaKantor    *string         `json:"noBaKantor" gorm:"type:varchar(30)"`
	TglBaKantor   *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan  *string         `json:"noBaLapangan" gorm:"type:varchar(30)"`
	TglBaLapangan *datatypes.Date `json:"tglBaLapangan"`
	TglCetak      *datatypes.Date `json:"tglCetak"`
	NipPenetak    *string         `json:"nipPenetak" gorm:"type:varchar(9)"`

	gormhelper.DateModel
}

func (i PstPermohonan) SetSkSk() *SkSk {
	return &SkSk{
		PermohonanId: &i.Id,
		KanwilId:     i.KanwilId,
		KppbbId:      i.KppbbId,
	}
}
