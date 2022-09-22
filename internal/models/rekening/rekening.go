package rekening

type Rekening struct {
	ID             uint64  `json:"id" gorm:"primaryKey"`
	ParentID       *uint64 `json:"parent_id"`
	Tipe           *string `json:"tipe"`
	Kelompok       *string `json:"kelompok"`
	Jenis          *string `json:"jenis"`
	Objek          *string `json:"objek"`
	Rincian        *string `json:"rincian"`
	Sub1           *string `json:"sub1"`
	Sub2           *string `json:"sub2"`
	Sub3           *string `json:"sub3"`
	Kode           *string `json:"kode"`
	Nama           *string `json:"nama"`
	Level          *int    `json:"level"`
	KodeBaru       *string `json:"kode_baru"`
	KodeJenisPajak *string `json:"kode_jenis_pajak"`
	KodeVaJatim    *string `json:"kode_va_jatim"`
	KodeBilling    *string `json:"kode_billing"`
	KodeJenisUsaha *string `json:"kode_jenis_usaha"`
	JenisUsaha     *string `json:"jenis_usaha"`
}
