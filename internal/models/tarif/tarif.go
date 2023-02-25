package tarif

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type Tarif struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode string  `json:"provinsi_kd" gorm:"type:varchar(2);not null"`
	Daerah_Kode   string  `json:"daerah_kd" gorm:"type:varchar(2);not null"`
	TahunAwal     string  `json:"tahunAwal" gorm:"size:4;not null"`
	TahunAkhir    string  `json:"tahunAkhir" gorm:"size:4"`
	NjopMin       float64 `json:"njopMin" gorm:"size:15"`
	NjopMax       float64 `json:"njopMax" gorm:"size:15"`
	NilaiTarif    float64 `json:"nilaiTarif" gorm:"type:decimal(3,3)"`
	gormhelper.DateModel
}

type CreateDto struct {
	Provinsi_Kode string  `json:"provinsi_kd" validate:"required"`
	Daerah_Kode   string  `json:"daerah_kd" validate:"required"`
	TahunAwal     string  `json:"tahunAwal" validate:"required"`
	TahunAkhir    string  `json:"tahunAkhir" validate:"required"`
	NjopMin       float64 `json:"njopMin" validate:"required"`
	NjopMax       float64 `json:"njopMax" validate:"required"`
	NilaiTarif    float64 `json:"nilaiTarif" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kd" gorm:"type:varchar(2);not null"`
	Daerah_Kode   *string  `json:"daerah_kd" gorm:"type:varchar(2);not null"`
	TahunAwal     *string  `json:"tahunAwal" gorm:"size:4;not null"`
	TahunAkhir    *string  `json:"tahunAkhir" gorm:"size:4"`
	NjopMin       *float64 `json:"njopMin" gorm:"size:15"`
	NjopMax       *float64 `json:"njopMax" gorm:"size:15"`
	NilaiTarif    *float64 `json:"nilaiTarif"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

type UpdateDto struct {
	TahunAwal  *string  `json:"tahunAwal"`
	TahunAkhir *string  `json:"tahunAkhir"`
	NjopMin    *float64 `json:"njopMin"`
	NjopMax    *float64 `json:"njopMax"`
	NilaiTarif *float64 `json:"nilaiTarif"`
}
