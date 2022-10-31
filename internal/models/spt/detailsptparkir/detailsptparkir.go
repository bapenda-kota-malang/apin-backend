package detailsptparkir

import mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"

type DetailSptParkir struct {
	Id             uint64            `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id         uint64            `json:"spt_id"`
	JenisKendaraan mt.JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      *uint64           `json:"kapasitas"`
}

type CreateDto struct {
	Spt_Id         uint
	JenisKendaraan mt.JenisKendaraan `json:"jenisKendaraan"  validate:"required"`
	Kapasitas      uint64            `json:"kapasitas"  validate:"required"`
}

type UpdateDto struct {
	Spt_Id         uint              `json:"spt_id"`
	JenisKendaraan mt.JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      uint64            `json:"kapasitas"`
}
