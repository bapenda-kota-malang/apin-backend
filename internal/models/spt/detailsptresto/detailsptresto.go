package detailsptresto

type DetailSptResto struct {
	Id               uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id           uint64   `json:"spt_id"`
	JumlahMeja       *uint64  `json:"jumlahMeja"`
	JumlahKursi      *uint64  `json:"jumlahKursi"`
	TarifMinuman     *float64 `json:"tarifMinuman" gorm:"type:decimal"`
	TarifMakanan     *float64 `json:"tarifMakanan" gorm:"type:decimal"`
	JumlahPengunjung *uint64  `json:"jumlahPengunjung"`
	IdBt             *uint64  `json:"idBt"`
}

type CreateDto struct {
	Spt_Id           uint
	JumlahMeja       uint64  `json:"jumlahMeja" validate:"required"`
	JumlahKursi      uint64  `json:"jumlahKursi" validate:"required"`
	TarifMinuman     float64 `json:"tarifMinuman" validate:"required"`
	TarifMakanan     float64 `json:"tarifMakanan" validate:"required"`
	JumlahPengunjung uint64  `json:"jumlahPengunjung" validate:"required"`
	IdBt             uint64  `json:"idBt" validate:"required"`
}

type UpdateDto struct {
	Spt_Id           uint    `json:"spt_id"`
	JumlahMeja       uint64  `json:"jumlahMeja"`
	JumlahKursi      uint64  `json:"jumlahKursi"`
	TarifMinuman     float64 `json:"tarifMinuman"`
	TarifMakanan     float64 `json:"tarifMakanan"`
	JumlahPengunjung uint64  `json:"jumlahPengunjung"`
	IdBt             uint64  `json:"idBt"`
}
