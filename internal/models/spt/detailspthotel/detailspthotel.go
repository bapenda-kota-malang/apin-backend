package detailspthotel

type DetailSptHotel struct {
	Id                  uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id              uint64   `json:"spt_id"`
	GolonganKamar       *string  `json:"golonganKamar" gorm:"type:varchar(20)"`
	Tarif               *float64 `json:"tarif" gorm:"type:decimal"`
	JumlahKamar         *uint64  `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint64  `json:"jumlahKamarYangLaku"`
}

type CreateDto struct {
	Spt_Id              uint
	GolonganKamar       string  `json:"golonganKamar" validate:"required"`
	Tarif               float64 `json:"tarif" validate:"required"`
	JumlahKamar         uint64  `json:"jumlahKamar" validate:"required"`
	JumlahKamarYangLaku uint64  `json:"jumlahKamarYangLaku" validate:"required"`
}

type UpdateDto struct {
	Spt_Id              uint    `json:"spt_id"`
	GolonganKamar       string  `json:"golonganKamar"`
	Tarif               float64 `json:"tarif"`
	JumlahKamar         uint64  `json:"jumlahKamar"`
	JumlahKamarYangLaku uint64  `json:"jumlahKamarYangLaku"`
}
