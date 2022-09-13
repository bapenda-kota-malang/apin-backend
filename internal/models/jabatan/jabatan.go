package jabatan

type Jabatan struct {
	Id   int    `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Kode string `json:"kode" gorm:"type:varchar"`
	Nama string `json:"nama" gorm:"type:varchar"`
}

type Create struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type GetList struct {
	Create
}

type GetDetail struct {
	Create
}

type Update struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type Filter struct {
	Create
}

func (j Jabatan) GetCreateStruct() any {
	return &Create{}
}

func (j Jabatan) GetListStruct() any {
	return &GetList{}
}

func (j Jabatan) GetDetailStruct() any {
	return &GetDetail{}
}

func (j Jabatan) GetUpdateStruct() any {
	return &Create{}
}

func (j Jabatan) GetFilterStruct() any {
	return &Filter{}
}
