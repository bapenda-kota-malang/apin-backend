package skpd

type SatuanKerja struct {
	Id     uint64  `json:"id" gorm:"primaryKey"`
	Kode   *string `json:"kode" gorm:"type:varchar(30)"`
	Nama   *string `json:"nama" gorm:"type:varchar(100)"`
	Alamat *string `json:"alamat" gorm:"type:varchar(255)"`
	Telp   *string `json:"telp" gorm:"type:varchar(20)"`
}

type CreateDto struct {
	Kode   string `json:"kode" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Telp   string `json:"telp" validate:"required"`
}

type UpdateDto struct {
	Kode   string `json:"kode"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Telp   string `json:"telp"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama"`
	Alamat   *string `json:"alamat"`
	Telp     *string `json:"telp"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}
