package configurationmodel

type (
	Skpd struct {
		ID       uint64  `json:"id" gorm:"primaryKey"`
		KodeSkpd *string `json:"kode_skpd"`
		Name     *string `json:"name"`
		Address  *string `json:"address"`
		Telp     *string `json:"telp"`
	}

	Rekening struct {
		ID               uint64  `json:"id" gorm:"primaryKey"`
		ParentID         *uint64 `json:"parent_id"`
		Tipe             *string `json:"tipe"`
		Kelompok         *string `json:"kelompok"`
		Jenis            *string `json:"jenis"`
		Objek            *string `json:"objek"`
		Rincian          *string `json:"rincian"`
		Sub1             *string `json:"sub1"`
		Sub2             *string `json:"sub2"`
		Sub3             *string `json:"sub3"`
		KodeRekening     *string `json:"kode_rekening"`
		Name             *string `json:"name"`
		Level            *int    `json:"level"`
		KodeRekeningBaru *string `json:"kode_rekening_baru"`
		KodeJenisPajak   *string `json:"kode_jenis_pajak"`
		KodeVaJatim      *string `json:"kode_va_jatim"`
		KodeBilling      *string `json:"kode_billing"`
		KodeJenisUsaha   *string `json:"kode_jenis_usaha"`
		JenisUsaha       *string `json:"jenis_usaha"`
	}

	VendorEtax struct{}

	Kecamatan struct {
		ID            uint64  `json:"id" gorm:"primaryKey"`
		KodeKecamatan *string `json:"kode_kecamatan"`
		Kecamatan     *string `json:"kecamatan"`
	}

	Kelurahan struct {
		ID            uint64  `json:"id" gorm:"primaryKey"`
		KodeKecamatan *string `json:"kode_kecamatan"`
		KodeKelurahan *string `json:"kode_kelurahan"`
		Kelurahan     *string `json:"kelurahan"`
	}
)
