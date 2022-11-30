package undanganpemeriksaan

type Status int
type StatusKirim int

const (
	StatusTerbit  Status = 1 // terbit
	StatusSelesai Status = 2 // selesai

	StatusTerkirim StatusKirim = 1 // terkirim
	StatusGagal    StatusKirim = 2 // gagal
)
