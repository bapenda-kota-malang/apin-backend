package configurationmodel

type (
	Skpd struct {
		ID     uint64  `json:"id" gorm:"primaryKey"`
		Kode   *string `json:"kode"`
		Nama   *string `json:"nama"`
		Alamat *string `json:"alamat"`
		Telp   *string `json:"telp"`
	}

	Rekening struct {
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

	VendorEtax struct{}

	Provinsi struct {
		ID     uint64    `json:"id" gorm:"primaryKey"`
		Kode   string    `json:"kode" gorm:"unique"`
		Nama   string    `json:"nama"`
		Daerah []*Daerah `json:"daerah,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
	}

	Daerah struct {
		ID           uint64       `json:"id" gorm:"primaryKey"`
		ProvinsiKode string       `json:"provinsi_kode"`
		Provinsi     *Provinsi    `json:"provinsi,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
		Kode         string       `json:"kode" gorm:"unique"`
		Nama         string       `json:"nama"`
		Kecamatan    []*Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
	}

	Kecamatan struct {
		ID         uint64       `json:"id" gorm:"primaryKey"`
		DaerahKode string       `json:"daerah_kode"`
		Daerah     *Daerah      `json:"daerah,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
		Kode       string       `json:"kode" gorm:"unique"`
		Nama       string       `json:"nama"`
		Kelurahan  []*Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
	}

	Kelurahan struct {
		ID            uint64     `json:"id" gorm:"primaryKey"`
		KecamatanKode string     `json:"kecamatan_kode"`
		Kecamatan     *Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
		Kode          string     `json:"kode" gorm:"unique"`
		Nama          string     `json:"nama"`
	}
)
