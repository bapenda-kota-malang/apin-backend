package registrasinpwpd

type DetailRegObjekPajak struct {
	Id                 uint64           `json:"id" gorm:"primaryKey"`
	RegistrasiNpwpd_Id uint64           `json:"registrasiNpwpd_id"`
	RegistrasiNpwpd    *RegistrasiNpwpd `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	JenisOp            *string          `json:"jenisOp" gorm:"size:200"`
	JumlahOp           *string          `json:"jumlahOp" gorm:"size:200"`
	TarifOp            *string          `json:"tarifOp" gorm:"size:200"`
	UnitOp             *string          `json:"unitOp" gorm:"size:50"`
	Notes              *string          `json:"notes" gorm:"size:200"`
}

type DetailRegObjekPajakAirTanah struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakHiburan struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakHotel struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakParkir struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakPpj struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakReklame struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakResto struct {
	DetailRegObjekPajak
}

type DetailRegObjekPajakCreate struct {
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}

type DetailRegObjekPajakUpdate struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	JumlahOp *string `json:"jumlahOp" gorm:"size:200"`
	TarifOp  *string `json:"tarifOp" gorm:"size:200"`
	UnitOp   *string `json:"unitOp" gorm:"size:50"`
	Notes    *string `json:"notes" gorm:"size:200"`
}
