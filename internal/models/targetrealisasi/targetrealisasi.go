package targetrealisasi

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type TargetRealisasi struct {
	Id              uint64   `json:"id" gorm:"primaryKey"`
	Tahun           *string  `json:"tahun" gorm:"type:char(4)"`
	JenisPajak_Kode *string  `json:"jenisPajak_kode" gorm:"type:varchar(10)"`
	Target          *float64 `json:"target"`
	Realisasi       *float64 `json:"realisasi"`
	gh.DateModel
}

type CreateDto struct {
	Tahun           string  `json:"tahun" validate:"required"`
	JenisPajak_Kode string  `json:"jenisPajak_kode" validate:"required"`
	Target          float64 `json:"target" validate:"required"`
}

type UpdateDto struct {
	Tahun           string  `json:"tahun" validate:"required"`
	JenisPajak_Kode string  `json:"jenisPajak_kode" validate:"required"`
	Target          float64 `json:"target" validate:"required"`
}

type FilterDto struct {
	Tahun           *string  `json:"tahun"`
	JenisPajak_Kode *string  `json:"jenisPajak_kode"`
	Target          *float64 `json:"target"`
}
