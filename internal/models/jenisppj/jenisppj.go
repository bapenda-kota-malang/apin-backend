package jenisppj

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type JenisPPJ struct {
	Id          uint    `json:"id" gorm:"primarykey"`
	Tahun       uint16  `json:"tahun"`
	Jenis       string  `json:"jenisPpj"`
	TarifPersen float32 `json:"tarif"`
	gormhelper.DateModel
}

type CreateDto struct {
	Tahun uint16  `json:"tahun"`
	Jenis string  `json:"jenisPpj"`
	Tarif float32 `json:"tarif"`
}

type UpdateDto struct {
	Tahun uint16  `json:"tahun"`
	Jenis string  `json:"jenisPpj"`
	Tarif float32 `json:"tarif"`
}

type FilterDto struct {
	Id    uint    `json:"id" gorm:"primarykey"`
	Tahun uint16  `json:"tahun"`
	Jenis string  `json:"jenisPpj"`
	Tarif float32 `json:"tarif"`
}
