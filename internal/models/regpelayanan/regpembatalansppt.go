package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type RegPembatalanSppt struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	PermohonanId *uint64 `json:"permohonanId" gorm:"type:integer"`
	PermohonanNOP
	TahunPajakSppt      *string         `json:"TahunPajakSppt" gorm:"type:varchar(4)"`
	PembatalanSpptKe    *string         `json:"pembatalanSpptKe" gorm:"type:varchar(2)"`
	KdBatal             *string         `json:"kdBatal" gorm:"type:varchar(2)"`
	KdKanwilBank        *string         `json:"kdKanwilBank" gorm:"type:varchar(2)"`
	KdKppbbBank         *string         `json:"kdKppbbBank" gorm:"type:varchar(2)"`
	KdBankTunggal       *string         `json:"kdBankTunggal" gorm:"type:varchar(2)"`
	KdBankPersepsi      *string         `json:"kdBankPersepsi" gorm:"type:varchar(2)"`
	KdTP                *string         `json:"kdTP" gorm:"type:varchar(2)"`
	DendaSppt           *int            `json:"dendaSppt" gorm:"type:INTEGER"`
	JmlSpptYgDibatalkan *int32          `json:"jmlSpptYgDibatalkan" gorm:"type:BIGINT"`
	TanggalPembatalan   *datatypes.Date `json:"tanggalPembatalan"`
	NipPembatal         *string         `json:"nipPembatal" gorm:"type:varchar(9)"`

	gormhelper.DateModel
}

func (i RegPstPermohonan) SetPembatalanSppt(nop PermohonanNOP) *RegPembatalanSppt {
	return &RegPembatalanSppt{
		PermohonanId:  &i.Id,
		PermohonanNOP: nop,
	}
}
