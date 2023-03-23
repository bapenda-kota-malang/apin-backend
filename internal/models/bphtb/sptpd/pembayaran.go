package bphtb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type PembayaranBphtb struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	Sptpd_Ref        *uuid.UUID      `json:"sptpd_id" gorm:"type:uuid"`
	NominalBayar     string          `json:"nominalBayar" gorm:"type:varchar(50)"`
	TglBayar         *datatypes.Date `json:"tglBayar"`
	PaymentPoint_Id  *string         `json:"paymentPoint_id"`
	KodeValidasiBank string          `json:"kodeValidasiBank" gorm:"type:varchar(100)"`
	CreatedBy        *uint64         `json:"createdBy"`
	BphtbSptpd       *BphtbSptpd     `json:"bphtbSptpd,omitempty" gorm:"foreignKey:Sptpd_Ref;references:Id"`

	gormhelper.DateModel
}

type RequestPembayaranBphtb struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	Sptpd_Ref        uuid.UUID       `json:"sptpd_id" gorm:"type:uuid"`
	NominalBayar     string          `json:"nominalBayar" gorm:"type:varchar(50)"`
	TglBayar         *datatypes.Date `json:"tglBayar"`
	PaymentPoint_Id  *string         `json:"paymentPoint_id"`
	KodeValidasiBank string          `json:"kodeValidasiBank" gorm:"type:varchar(100)"`
	CreatedBy        *uint64         `json:"createdBy"`
	gormhelper.DateModel
}

type ResponsePembayaranBphtb struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	Sptpd_Ref        uuid.UUID       `json:"sptpd_id" gorm:"type:uuid"`
	NominalBayar     string          `json:"nominalBayar" gorm:"type:varchar(50)"`
	TglBayar         *datatypes.Date `json:"tglBayar"`
	PaymentPoint_Id  *string         `json:"paymentPoint_id"`
	KodeValidasiBank string          `json:"kodeValidasiBank" gorm:"type:varchar(100)"`
	CreatedBy        *uint64         `json:"createdBy"`
	BphtbSptpd       *BphtbSptpd     `json:"bphtbSptpd,omitempty"`
	Page             int             `json:"page"`
	PageSize         int             `json:"page_size"`
	gormhelper.DateModel
}
