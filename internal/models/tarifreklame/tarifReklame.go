package tarifreklame

import "github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"

type JenisMasa int16

const (
	MasaPajakTahun         JenisMasa = 1 //masa pajak tetap 1 tahun
	MasaPajakBulan         JenisMasa = 2 //masa pajak insidentil 1 bulan
	MasaPajakHari          JenisMasa = 3 //masa pajak insidentil 1 hari
	MasaPajakPenyelenggara JenisMasa = 4 //masa pajak insidentil 1 kali penyelenggaraan
)

type TarifReklame struct {
	Id                  uint64                             `json:"id" gorm:"primaryKey"`
	JenisMasa           JenisMasa                          `json:"jenisMasa" gorm:"type:integer"`
	JenisReklame        *string                            `json:"jenisReklame" gorm:"type:varchar(200)"`
	DasarPengenaan      *string                            `json:"dasarPengenaan" gorm:"type:varchar(100)"`
	KlasifikasiJalan_Id *string                            `json:"klasifikasiJalan_id" gorm:"type:varchar(3)"`
	KlasifikasiJalan    *klasifikasijalan.KlasifikasiJalan `json:"omitempty" gorm:"foreignKey:KlasifikasiJalan_Id"`
	MasaPajak           *string                            `json:"masaPajak,omitempty" gorm:"size:100"`
	Tarif               *float64                           `json:"tarif" gorm:"type:decimal"`
}

type CreateDto struct {
	JenisMasa           JenisMasa `json:"jenisMasa" validate:"required"`
	JenisReklame        string    `json:"jenisReklame" validate:"required"`
	DasarPengenaan      string    `json:"dasarPengenaan" validate:"required"`
	KlasifikasiJalan_Id string    `json:"klasifikasiJalan_id" validate:"required"`
	MasaPajak           *string   `json:"masaPajak,omitempty"`
	Tarif               float64   `json:"tarif" validate:"required"`
}

type UpdateDto struct {
	JenisMasa           JenisMasa `json:"jenisMasa"`
	JenisReklame        *string   `json:"jenisReklame"`
	DasarPengenaan      string    `json:"dasarPengenaan"`
	KlasifikasiJalan_Id string    `json:"klasifikasiJalan_id"`
	MasaPajak           *string   `json:"masaPajak,omitempty"`
	Tarif               float64   `json:"tarif"`
}

type FilterDto struct {
	// JenisMasa           JenisMasa `json:"jenisMasa"`
	JenisReklame        *string  `json:"jenisReklame"`
	DasarPengenaan      *string  `json:"dasarPengenaan"`
	KlasifikasiJalan_Id *string  `json:"klasifikasiJalan_id"`
	Tarif               *float64 `json:"tarif"`
	MasaPajak           *string  `json:"masaPajak,omitempty"`
	Page                int      `json:"page"`
	PageSize            int      `json:"page_size"`
}
