package detailesptair

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DetailEspt struct {
	Id      uint `json:"id" gorm:"primarykey"`
	Espt_Id uint `json:"espt_id"`
}

type DetailEsptAir struct {
	DetailEspt
	Peruntukan string  `json:"peruntukan" gorm:"size:100"`
	JenisAbt   string  `json:"jenisAbt" gorm:"20"`
	Pengenaan  float32 `json:"pengenaan"`
	gormhelper.DateModel
}
