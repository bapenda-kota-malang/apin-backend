package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type RegSPMKP struct {
	Id              uint64          `json:"id" gorm:"primaryKey"`
	PermohonanId    *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId        *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId         *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	JnsSK           *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSpmkp         *string         `json:"noSpmkp" gorm:"type:varchar(30)"`
	TglSpmkp        *datatypes.Date `json:"tglSpmkp"`
	TahunPelayanan  *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan *string         `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanNOP

	NoRekWp  *string `json:"noRekWp" gorm:"type:varchar(20)"`
	NmBankWp *string `json:"nmBankWp" gorm:"type:varchar(30)"`
	NipRekam *string `json:"nipRekam" gorm:"type:varchar(9)"`

	gormhelper.DateModel
}

func (i RegPstPermohonan) SetSPMKP(nop PermohonanNOP) *RegSPMKP {
	return &RegSPMKP{
		PermohonanId:    &i.Id,
		KanwilId:        i.KanwilId,
		KppbbId:         i.KppbbId,
		TahunPelayanan:  i.TahunPelayanan,
		BundelPelayanan: i.BundelPelayanan,
		NoUrutPelayanan: i.NoPelayanan,
		PermohonanNOP:   nop}
}
