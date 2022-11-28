package kanwil

type Kanwil struct {
	Id               uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	AlKanwil         *string `json:"alKanwil" gorm:"type:varchar(50)"`
	KdKanwil         *string `json:"kdKanwil,omitempty" gorm:"type:varchar(2)"`
	KotaTerbitKanwil *string `json:"kotaTerbitKanwil" gorm:"type:varchar(30)"`
	NmKanwil         *string `json:"nmKanwil" gorm:"type:varchar(30)"`
	NoFaksimili      *string `json:"noFaksimili" gorm:"type:varchar(50)"`
	NoTelpon         *string `json:"noTelpon" gorm:"type:varchar(50)"`
}
