package npwpd

type DetailObjekPajak struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	Npwpd_Id uint64  `json:"npwpd_id"`
	Npwpd    *Npwpd  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	JenisOp  *string `json:"jenisOp" gorm:"size:200"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type DetailObjekPajakAirTanah struct {
	DetailObjekPajak
}

type DetailObjekPajakHiburan struct {
	DetailObjekPajak
}

type DetailObjekPajakHotel struct {
	DetailObjekPajak
}

type DetailObjekPajakParkir struct {
	DetailObjekPajak
}

type DetailObjekPajakPpj struct {
	DetailObjekPajak
}

type DetailObjekPajakReklame struct {
	DetailObjekPajak
}

type DetailObjekPajakResto struct {
	DetailObjekPajak
}
type DetailOpCreate struct {
	Npwpd_Id uint64  `json:"npwpd_id"`
	JenisOp  *string `json:"jenisOp" gorm:"size:200"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type DetailOpUpdate struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}
