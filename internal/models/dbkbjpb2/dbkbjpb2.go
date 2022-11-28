package dbkbjpb2

type DbkbJpb2 struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Id   *uint64 `json:"provinsi_id"`
	Daerah_Id     *uint64 `json:"daerah_id"`
	TahunDbkbJpb2 *string `json:"tahunDbkbJpb2" gorm:"type:char(4)"`
}
