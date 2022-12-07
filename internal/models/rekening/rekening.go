package rekening

type Rekening struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Parent_Id      *string `json:"parent_id" gorm:"type:varchar(5)"`
	Tipe           *string `json:"tipe" gorm:"type:varchar(5)"`
	Kelompok       *string `json:"kelompok" gorm:"type:varchar(5)"`
	Jenis          *string `json:"jenis" gorm:"type:varchar(5)"`
	Objek          *string `json:"objek" gorm:"type:varchar(5)"`
	Rincian        *string `json:"rincian" gorm:"type:varchar(5)"`
	Sub1           *string `json:"sub1" gorm:"type:varchar(5)"`
	Sub2           *string `json:"sub2" gorm:"type:varchar(5)"`
	Sub3           *string `json:"sub3" gorm:"type:varchar(5)"`
	Kode           *string `json:"kode" gorm:"type:varchar(50)"`
	Nama           *string `json:"nama" gorm:"type:varchar(200)"`
	Level          *int    `json:"level" gorm:"type:integer"`
	KodeBaru       *string `json:"kodeBaru" gorm:"type:varchar(7)"`
	KodeJenisPajak *string `json:"kodeJenisPajak" gorm:"type:varchar(5)"`
	KodeVaJatim    *string `json:"kodeVaJatim" gorm:"type:varchar(100)"`
	KodeBilling    *string `json:"kodeBilling" gorm:"type:varchar(20)"`
	KodeJenisUsaha *string `json:"kodeJenisUsaha" gorm:"type:varchar(5)"`
	JenisUsaha     *string `json:"jenisUsaha" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	Parent_Id      string `json:"parent_id"`
	Tipe           string `json:"tipe"`
	Kelompok       string `json:"kelompok"`
	Jenis          string `json:"jenis"`
	Objek          string `json:"objek"`
	Rincian        string `json:"rincian"`
	Sub1           string `json:"sub1"`
	Sub2           string `json:"sub2"`
	Sub3           string `json:"sub3"`
	Kode           string `json:"kode"`
	Nama           string `json:"nama"`
	Level          int    `json:"level"`
	KodeBaru       string `json:"kodeBaru"`
	KodeJenisPajak string `json:"kodeJenisPajak"`
	KodeVaJatim    string `json:"kodeVaJatim"`
	KodeBilling    string `json:"kodeBilling"`
	KodeJenisUsaha string `json:"kodeJenisUsaha"`
	JenisUsaha     string `json:"jenisUsaha"`
}
type UpdateDto struct {
	Parent_Id      string `json:"parent_id"`
	Tipe           string `json:"tipe"`
	Kelompok       string `json:"kelompok"`
	Jenis          string `json:"jenis"`
	Objek          string `json:"objek"`
	Rincian        string `json:"rincian"`
	Sub1           string `json:"sub1"`
	Sub2           string `json:"sub2"`
	Sub3           string `json:"sub3"`
	Kode           string `json:"kode"`
	Nama           string `json:"nama"`
	Level          int    `json:"level"`
	KodeBaru       string `json:"kodeBaru"`
	KodeJenisPajak string `json:"kodeJenisPajak"`
	KodeVaJatim    string `json:"kodeVaJatim"`
	KodeBilling    string `json:"kodeBilling"`
	KodeJenisUsaha string `json:"kodeJenisUsaha"`
	JenisUsaha     string `json:"jenisUsaha"`
}
type FilterDto struct {
	Parent_Id          *string `json:"parent_id"`
	Tipe               *string `json:"tipe"`
	Kelompok           *string `json:"kelompok"`
	Jenis              *string `json:"jenis"`
	Objek              *string `json:"objek"`
	Rincian            *string `json:"rincian"`
	Sub1               *string `json:"sub1"`
	Sub2               *string `json:"sub2"`
	Sub3               *string `json:"sub3"`
	Kode               *string `json:"kode"`
	Nama               *string `json:"nama"`
	Level              *int    `json:"level"`
	KodeBaru           *string `json:"kodeBaru"`
	KodeJenisPajak     *string `json:"kodeJenisPajak"`
	KodeVaJatim        *string `json:"kodeVaJatim"`
	KodeBilling        *string `json:"kodeBilling"`
	KodeJenisUsaha     *string `json:"kodeJenisUsaha"`
	KodeJenisUsaha_Opt *string `json:"kodeJenisUsaha_opt"`
	JenisUsaha         *string `json:"jenisUsaha"`
	Page               int     `json:"page"`
	PageSize           int     `json:"page_size"`
	NoPagination       bool    `json:"no_pagination"`
}
