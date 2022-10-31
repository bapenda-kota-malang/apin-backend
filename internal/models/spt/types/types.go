package spt

type SptStatus string
type JenisKendaraan int16
type Peruntukan string
type JenisAbt string
type JenisReklame uint16

const (
	StatusBelumLunas          SptStatus = "00" //belum lunas
	StatusLunas               SptStatus = "11" //lunas penuh
	StatusKurangBayar         SptStatus = "12" //kurang bayar
	StatusKurangBayarAngsuran SptStatus = "13" //kurang bayar angsuran
	StatusLebihBayar          SptStatus = "14" //lebih bayar
	StatusSalahPenetapan      SptStatus = "21" //salah penetapan
	StatusDouble              SptStatus = "22" //double
	StatusRestitusiPenuh      SptStatus = "31" //restitusi penuh
	StatusRestitusiSebagian   SptStatus = "32" //restitusi sebagian

	KendaraanRodaDua   JenisKendaraan = 1 //roda dua
	KendaraanRodaEmpat JenisKendaraan = 2 //roda empat

	PeruntukanNonNiaga Peruntukan = "1" //non niaga
	PeruntukanNiaga    Peruntukan = "2" //niaga
	PeruntukanIndustri Peruntukan = "3" //industri dengan bahan baku
	PeruntukanPdam     Peruntukan = "4" //pdam

	AbtMataAir      JenisAbt = "1" //mata air
	AbtBukanMataAir JenisAbt = "2" //bukan mata air

	ReklameInsidentil JenisReklame = 1 //reklame insidentil
	ReklameTetap      JenisReklame = 2 //reklame tetap
)

type DetailSpt struct {
	Id     uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id uint64 `json:"spt_id"`
	// Spt    *ms.Spt `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
}
