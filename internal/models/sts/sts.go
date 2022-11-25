package sts

import (
	"time"

	mj "github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	mp "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
)

type Sts struct {
	Id                           uint64             `json:"id" gorm:"primaryKey"`
	NomorUrut                    int                `json:"nomorUrut"`
	NomorOutput                  string             `json:"nomorOutput"`
	TanggalSetor                 *time.Time         `json:"tanggalSetor"`
	Jurnal_Id                    *uint64            `json:"jurnal_id"`
	Jurnal                       *mj.Jurnal         `json:"jurnal" gorm:"foreignKey:Jurnal_Id"`
	AkunBendahara_Rekening_Id    *uint64            `json:"akunBendahara_rekening_id"`
	Rekening                     *rekening.Rekening `json:"rekening" gorm:"foreignKey:AkunBendahara_Rekening_Id"`
	BendaharaPenerima_Pegawai_Id *uint64            `json:"bendaharaPenerima_pegawai_id"`
	BendaharaPenerima            *mp.Pegawai        `json:"bendaharaPenerima,omitempty" gorm:"foreignKey:BendaharaPenerima_Pegawai_Id"`
	IdBt                         *int               `json:"idBt"`
	Keterangan                   *string            `json:"keterangan" gorm:"type:varchar(255)"`
	TotalSetor                   *float64           `json:"totalSetor" gorm:"type:decimal"`
	Id_Aktivitas                 *int               `json:"id_aktivitas"`
	Name                         *string            `json:"name" gorm:"type:varchar(100)"`
	TanggalSts                   *time.Time         `json:"tanggalSts"`
	IsSetor                      bool               `json:"isSetor"`
	StsDetails                   *[]StsDetail       `json:"stsDetail,omitempty" gorm:"foreignKey:Sts_Id;references:Id"`
	SumberDanaStss               *[]SumberDanaSts   `json:"sumberDanaSts,omitempty" gorm:"foreignKey:Sts_Id;references:Id"`
	SspdDetails                  *[]ms.SspdDetail   `json:"sspdDetail,omitempty" gorm:"foreignKey:Sts_Id;references:Id"`
}

type CreateDto struct {
	TanggalSetor                 *string                   `json:"tanggalSetor"`
	Jurnal_Id                    *uint64                   `json:"jurnal_id"`
	AkunBendahara_Rekening_Id    *uint64                   `json:"akunBendahara_rekening_id"`
	BendaharaPenerima_Pegawai_Id *uint64                   `json:"bendaharaPenerima_pegawai_id"`
	IdBt                         *int                      `json:"idBt"`
	Keterangan                   *string                   `json:"keterangan"`
	TotalSetor                   *float64                  `json:"totalSetor"`
	Id_Aktivitas                 *int                      `json:"id_aktivitas"`
	Name                         *string                   `json:"name"`
	TanggalSts                   *string                   `json:"tanggalSts"`
	IsSetor                      bool                      `json:"isSetor"`
	StsDetail                    *[]StsDetailCreateDto     `json:"stsDetail"`
	SumberDanaSts                *[]SumberDanaStsCreateDto `json:"sumberDanaSts"`
}
type UpdateDto struct {
	Id                           *uint64                   `json:"id"`
	TanggalSetor                 *string                   `json:"tanggalSetor"`
	Jurnal_Id                    *uint64                   `json:"jurnal_id"`
	AkunBendahara_Rekening_Id    *uint64                   `json:"akunBendahara_rekening_id"`
	BendaharaPenerima_Pegawai_Id *uint64                   `json:"bendaharaPenerima_pegawai_id"`
	IdBt                         *int                      `json:"idBt"`
	Keterangan                   *string                   `json:"keterangan"`
	TotalSetor                   *float64                  `json:"totalSetor"`
	Id_Aktivitas                 *int                      `json:"id_aktivitas"`
	Name                         *string                   `json:"name"`
	TanggalSts                   *string                   `json:"tanggalSts"`
	IsSetor                      bool                      `json:"isSetor"`
	StsDetail                    *[]StsDetailUpdateDto     `json:"stsDetail"`
	SumberDanaSts                *[]SumberDanaStsUpdateDto `json:"sumberDanaSts"`
}

type FilterDto struct {
	Id                           *uint64    `json:"id"`
	StsNumber                    *int       `json:"stsNumber"`
	TanggalSetor                 *time.Time `json:"tanggalSetor"`
	Jurnal_Id                    *uint64    `json:"jurnal_id"`
	AkunBendahara_Rekening_Id    *uint64    `json:"akunBendahara_rekening_id"`
	BendaharaPenerima_Pegawai_Id *uint64    `json:"bendaharaPenerima_pegawai_id"`
	IdBt                         *int       `json:"idBt"`
	Keterangan                   *string    `json:"keterangan"`
	TotalSetor                   *float64   `json:"totalSetor"`
	Id_Aktivitas                 *int       `json:"id_aktivitas"`
	Name                         *string    `json:"name"`
	TanggalSts                   *time.Time `json:"tanggalSts"`
	IsSetor                      *bool      `json:"isSetor"`
	Page                         int        `json:"page"`
	PageSize                     int        `json:"page_size"`
}
