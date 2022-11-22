package jenisusaha

import "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"

type JenisUsahaDetail struct {
	Id            uint64             `json:"id" gorm:"primaryKey"`
	JenisUsaha_Id uint64             `json:"jenisUsaha_id"`
	JenisUsaha    *JenisUsaha        `json:"jenisUsaha,omitempty" gorm:"foreignKey:JenisUsaha_Id"`
	Rekening_Id   *uint64            `json:"rekening_id"`
	Rekening      *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
}

type JenisUsahaDetailCreateDto struct {
	JenisUsaha_Id uint64  `json:"jenisUsaha_id"`
	Rekening_Id   *uint64 `json:"rekening_id"`
}

type JenisUsahaDetailUpdateDto struct {
	Id            *uint64 `json:"id"`
	JenisUsaha_Id *uint64 `json:"jenisUsaha_id"`
	Rekening_Id   *uint64 `json:"rekening_id"`
}
