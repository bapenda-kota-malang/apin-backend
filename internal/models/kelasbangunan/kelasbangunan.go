package kelasbangunan

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type KelasBangunan struct {
	Id                      uint64   `json:"id" gorm:"primaryKey"`
	KdBangunan              string   `json:"kdBangunan" gorm:"type:varchar(3)"`
	NilaiMaxBangunan        *float64 `json:"nilaiMaxBangunan" gorm:"type:decimal(8,2)"`
	NilaiMinBangunan        *float64 `json:"nilaiMinBangunan" gorm:"type:decimal(8,2)"`
	NilaiPerM2Bangunan      *float64 `json:"nilaiPerM2Bangunan" gorm:"type:decimal(8,2)"`
	TahunAkhirKelasBangunan *string  `json:"tahunAkhirKelasBangunan" gorm:"type:char(4)"`
	TahunAwalKelasBangunan  *string  `json:"tahunAwalKelasBangunan" gorm:"type:char(4)"`
	gh.DateModel
}

type CreateDto struct {
	KdBangunan              string   `json:"kdBangunan"`
	NilaiMaxBangunan        *float64 `json:"nilaiMaxBangunan"`
	NilaiMinBangunan        *float64 `json:"nilaiMinBangunan"`
	NilaiPerM2Bangunan      *float64 `json:"nilaiPerM2Bangunan"`
	TahunAkhirKelasBangunan *string  `json:"tahunAkhirKelasBangunan"`
	TahunAwalKelasBangunan  *string  `json:"tahunAwalKelasBangunan"`
}

type UpdateDto struct {
	KdBangunan              string   `json:"kdBangunan"`
	NilaiMaxBangunan        *float64 `json:"nilaiMaxBangunan"`
	NilaiMinBangunan        *float64 `json:"nilaiMinBangunan"`
	NilaiPerM2Bangunan      *float64 `json:"nilaiPerM2Bangunan"`
	TahunAkhirKelasBangunan *string  `json:"tahunAkhirKelasBangunan"`
	TahunAwalKelasBangunan  *string  `json:"tahunAwalKelasBangunan"`
}

type FilterDto struct {
	KdBangunan              string   `json:"kdBangunan"`
	NilaiMaxBangunan        *float64 `json:"nilaiMaxBangunan"`
	NilaiMinBangunan        *float64 `json:"nilaiMinBangunan"`
	NilaiPerM2Bangunan      *float64 `json:"nilaiPerM2Bangunan"`
	TahunAkhirKelasBangunan *string  `json:"tahunAkhirKelasBangunan"`
	TahunAwalKelasBangunan  *string  `json:"tahunAwalKelasBangunan"`
	Page                    int      `json:"page"`
	PageSize                int      `json:"page_size"`
}
