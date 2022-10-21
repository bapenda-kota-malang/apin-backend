package jenisppj

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type JenisPPJ struct {
	Id    uint    `json:"id" gorm:"primarykey"`
	Tahun uint16  `json:"tahun"`
	Jenis string  `json:"jenisPpj"`
	Tarif float32 `json:"tarif"`
	gormhelper.DateModel
}
