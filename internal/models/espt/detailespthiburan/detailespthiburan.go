package detailespthiburan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/lib/pq"
)

type DetailEsptHiburan struct {
	detailesptair.DetailEspt
	Kelas              pq.StringArray `json:"kelas" gorm:"type:varchar[]"`
	Tarif              pq.Int64Array  `json:"tarif" gorm:"type:integer[]"`
	PengunjungWeekday  *uint          `json:"pengunjungWeekday" gorm:"comment:khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL"`
	PengunjungWeekend  *uint          `json:"pengunjungWeekend" gorm:"comment:khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL"`
	PertunjukanWeekday *uint          `json:"pertunjukanWeekday" gorm:"comment:khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL"`
	PertunjukanWeekend *uint          `json:"pertunjukanWeekend" gorm:"comment:khusus untuk pertunjukan Film, Kesenian dan Sejenisnya, Pagelaran Musik dan Tari. NULL"`
	JumlahMeja         *uint          `json:"jumlahMeja" gorm:"comment:Khusus untuk Billyar, Permainan Ketangkasan. NULL"`
	JumlahRuangan      *int           `json:"jumlahRuangan" gorm:"comment:khusus untuk Panti Pijat, Mandi Uap/Spa, Karaoke. NULL"`
	KarcisBebas        bool           `json:"karcisBebas"`
	JumlahKarcisBebas  *uint32        `json:"jumlahKarcisBebas" gorm:"comment:kalo KarcisBebas Iya, ini diisi"`
	MesinTiket         bool           `json:"mesinTiket"`
	Pembukuan          bool           `json:"pembukuan"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id            uint
	Kelas              pq.StringArray `json:"kelas"`
	Tarif              pq.Int64Array  `json:"tarif"`
	PengunjungWeekday  *uint          `json:"pengunjungWeekday"`
	PengunjungWeekend  *uint          `json:"pengunjungWeekend"`
	PertunjukanWeekday *uint          `json:"pertunjukanWeekday"`
	PertunjukanWeekend *uint          `json:"pertunjukanWeekend"`
	JumlahMeja         *uint          `json:"jumlahMeja"`
	JumlahRuangan      *uint          `json:"jumlahRuangan"`
	KarcisBebas        bool           `json:"karcisBebas"`
	JumlahKarcisBebas  *uint32        `json:"jumlahKarcisBebas" validate:"min=0"`
	MesinTiket         bool           `json:"mesinTiket"`
	Pembukuan          bool           `json:"pembukuan"`
}

type UpdateDto struct {
	Id                 uint           `json:"id"`
	Kelas              pq.StringArray `json:"kelas"`
	Tarif              pq.Int64Array  `json:"tarif"`
	PengunjungWeekday  *uint          `json:"pengunjungWeekday"`
	PengunjungWeekend  *uint          `json:"pengunjungWeekend"`
	PertunjukanWeekday *uint          `json:"pertunjukanWeekday"`
	PertunjukanWeekend *uint          `json:"pertunjukanWeekend"`
	JumlahMeja         *uint          `json:"jumlahMeja"`
	JumlahRuangan      *uint          `json:"jumlahRuangan"`
}
