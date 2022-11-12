package sts

import (
	"time"

	mj "github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
)

type Sts struct {
	Id                        uint64             `json:"id" gorm:"primaryKey"`
	StsNumber                 *string            `json:"stsNumber" gorm:"type:varchar(50)"`
	TanggalSetor              *time.Time         `json:"tanggalSetor"`
	Jurnal_Id                 *uint64            `json:"jurnal_id"`
	Jurnal                    *mj.Jurnal         `json:"jurnal" gorm:"foreignKey:Jurnal_Id"`
	AkunBendahara_Rekening_Id *uint64            `json:"akunBendahara_rekening_id"`
	Rekening                  *rekening.Rekening `json:"rekening" gorm:"foreignKey:AkunBendahara_Rekening_Id"`
	IdBt                      *int               `json:"idBt"`
	Keterangan                *string            `json:"keterangan" gorm:"type:varchar(255)"`
	TotalSetor                *float64           `json:"totalSetor" gorm:"type:decimal"`
	Id_Aktivitas              *int               `json:"id_aktivitas"`
	Name                      *string            `json:"name" gorm:"type:varchar(100)"`
	TanggalSts                *time.Time         `json:"tanggalSts"`
	IsSetor                   *bool              `json:"isSetor"`
}
