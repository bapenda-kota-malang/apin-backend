package bankpersepsi

type BankPersepsi struct {
	Id          uint64 `json:"id" gorm:"primaryKey"`
	Kanwil_Kode string `json:"kanwil_kode" gorm:"type:char(2)"`
	KPPBB_Kode  string `json:"kppbb_kode" gorm:"type:char(2)"`
	Kode        string `json:"kode" gorm:"unique;type:char(2)"`
	Nama        string `json:"nama" gorm:"type:varchar(30)"`
	Alamat      string `json:"alamat" gorm:"type:varchar(50)"`
	NoRek       string `json:"norek" gorm:"type:varchar(15)"`
}

type CreateDto struct {
	Kanwil_Kode string `json:"kanwil_kode"`
	KPPBB_Kode  string `json:"kppbb_kode"`
	Kode        string `json:"kode"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	NoRek       string `json:"norek"`
}

type UpdateDto struct {
	Id          *uint64 `json:"id"`
	Kanwil_Kode string  `json:"kanwil_kode"`
	KPPBB_Kode  string  `json:"kppbb_kode"`
	Kode        string  `json:"kode"`
	Nama        string  `json:"nama"`
	Alamat      string  `json:"alamat"`
	NoRek       string  `json:"norek"`
}

type FilterDto struct {
	Kanwil_Kode string `json:"kanwil_kode"`
	KPPBB_Kode  string `json:"kppbb_kode"`
	Kode        string `json:"kode"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	NoRek       string `json:"norek"`
	Page        int    `json:"page"`
	PageSize    int    `json:"page_size"`
}
