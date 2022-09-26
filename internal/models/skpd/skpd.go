package skpd

type Skpd struct {
	ID     uint64  `json:"id" gorm:"primaryKey"`
	Kode   *string `json:"kode"`
	Nama   *string `json:"nama"`
	Alamat *string `json:"alamat"`
	Telp   *string `json:"telp"`
}
