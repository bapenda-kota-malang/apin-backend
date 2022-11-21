package permohonan

import "time"

type Permohonan struct {
	Id                uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	StatusKolektif    *string   `json:"statusKolektif" gorm:"type:varchar(5)"`
	NoSuratPermohonan *string   `json:"noSuratPermohonan" gorm:"type:varchar(50)"`
	JenisPelayanan    *string   `json:"jenisPelayanan" gorm:"type:varchar(5)"`
	TanggalTerima     time.Time `json:"tanggalTerima"`
	TanggalSelesai    time.Time `json:"tanggalSelesai"`
	TanggalPermohonan time.Time `json:"tanggalPermohonan"`
	NOP               *string   `json:"nop" gorm:"type:varchar(50)"`
	NamaWP            *string   `json:"namaWP" gorm:"type:varchar(50)"`
	LetakOP           *string   `json:"letakOP" gorm:"type:varchar(255)"`
	Keterangan        *string   `json:"keterangan" gorm:"type:varchar(255)"`
	TahunPajak        *string   `json:"tahunPajak" gorm:"type:varchar(4)"`
	PenerimaanBerkas  *string   `json:"penerimaanBerkas" gorm:"type:varchar(255)"`
	Catatan           *string   `json:"catatan" gorm:"type:varchar(255)"`
}

type CreateDto struct {
	StatusKolektif    *string `json:"statusKolektif"`
	NoSuratPermohonan *string `json:"noSuratPermohonan"`
	JenisPelayanan    *string `json:"jenisPelayanan"`
	TanggalTerima     *string `json:"tanggalTerima"`
	TanggalSelesai    *string `json:"tanggalSelesai"`
	TanggalPermohonan *string `json:"tanggalPermohonan"`
	NOP               *string `json:"nop"`
	NamaWP            *string `json:"namaWP"`
	LetakOP           *string `json:"letakOP"`
	Keterangan        *string `json:"keterangan"`
	TahunPajak        *string `json:"tahunPajak"`
	PenerimaanBerkas  *string `json:"penerimaanBerkas"`
	Catatan           *string `json:"catatan"`
}

type UpdateDto struct {
	StatusKolektif    *string `json:"statusKolektif"`
	NoSuratPermohonan *string `json:"noSuratPermohonan"`
	JenisPelayanan    *string `json:"jenisPelayanan"`
	TanggalTerima     *string `json:"tanggalTerima"`
	TanggalSelesai    *string `json:"tanggalSelesai"`
	TanggalPermohonan *string `json:"tanggalPermohonan"`
	NOP               *string `json:"nop"`
	NamaWP            *string `json:"namaWP"`
	LetakOP           *string `json:"letakOP"`
	Keterangan        *string `json:"keterangan"`
	TahunPajak        *string `json:"tahunPajak"`
	PenerimaanBerkas  *string `json:"penerimaanBerkas"`
	Catatan           *string `json:"catatan"`
}

type FilterDto struct {
	StatusKolektif    *string `json:"statusKolektif"`
	NoSuratPermohonan *string `json:"noSuratPermohonan"`
	JenisPelayanan    *string `json:"jenisPelayanan"`
	TanggalTerima     *string `json:"tanggalTerima"`
	TanggalSelesai    *string `json:"tanggalSelesai"`
	TanggalPermohonan *string `json:"tanggalPermohonan"`
	NOP               *string `json:"nop"`
	NamaWP            *string `json:"namaWP"`
	LetakOP           *string `json:"letakOP"`
	TahunPajak        *string `json:"tahunPajak"`
	PenerimaanBerkas  *string `json:"penerimaanBerkas"`
}
