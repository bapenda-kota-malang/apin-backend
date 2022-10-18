package detailespthiburan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetaileEsptHiburan struct {
	detailesptair.DetailEspt
	Kelas              uint  `json:"kelas"`
	Tarif              uint  `json:"tarif"`
	PengunjungWeekday  *uint `json:"pengunjungWeekday" gorm:"comment:'khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL'"`
	PengunjungWeekend  *uint `json:"pengunjungWeekend" gorm:"comment:'khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL'"`
	PertunjukanWeekday *uint `json:"pertunjukanWeekday" gorm:"comment:'khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL'"`
	PertunjukanWeekend *uint `json:"pertunjukanWeekend" gorm:"comment:'khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL'"`
	JumlahMeja         *uint `json:"jumlahMeja" gorm:"comment:'Khusus untuk Billyar, Permainan Ketangkasan. NULL'"`
	JumlahRuangan      *uint `json:"jumlahRuangan" gorm:"comment:'khusus untuk Panti Pijat, Mandi Uap/Spa, Karaoke. NULL'"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id            *uint
	Kelas              uint  `json:"kelas"`
	Tarif              uint  `json:"tarif"`
	PengunjungWeekday  *uint `json:"pengunjungWeekday"`
	PengunjungWeekend  *uint `json:"pengunjungWeekend"`
	PertunjukanWeekday *uint `json:"pertunjukanWeekday"`
	PertunjukanWeekend *uint `json:"pertunjukanWeekend"`
	JumlahMeja         *uint `json:"jumlahMeja"`
	JumlahRuangan      *uint `json:"jumlahRuangan"`
}

type UpdateDto struct {
	Id                 uint  `json:"id"`
	Kelas              uint  `json:"kelas"`
	Tarif              uint  `json:"tarif"`
	PengunjungWeekday  *uint `json:"pengunjungWeekday"`
	PengunjungWeekend  *uint `json:"pengunjungWeekend"`
	PertunjukanWeekday *uint `json:"pertunjukanWeekday"`
	PertunjukanWeekend *uint `json:"pertunjukanWeekend"`
	JumlahMeja         *uint `json:"jumlahMeja"`
	JumlahRuangan      *uint `json:"jumlahRuangan"`
}
