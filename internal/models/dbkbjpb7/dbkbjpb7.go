package dbkbjpb7

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb7 struct {
	Id              uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode   *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode     *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb7   *string  `json:"tahunDbkbJpb7" gorm:"type:char(4)"`
	JenisDbkbJpb7   *string  `json:"jenisDbkbJpb7" gorm:"type:char(1)"`
	BintangDbkbJpb7 *string  `json:"bintangDbkbJpb7" gorm:"type:char(1)"`
	LantaiMinJpb7   *int     `json:"lantaiMinJpb7"`
	LantaiMaxJpb7   *int     `json:"lantaiMaxJpb7"`
	NilaiDbkbJp75   *float64 `json:"nilaiDbkbJpb7"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode   *string  `json:"provinsi_kode"`
	Daerah_Kode     *string  `json:"daerah_kode"`
	TahunDbkbJpb7   *string  `json:"tahunDbkbJpb7"`
	JenisDbkbJpb7   *string  `json:"jenisDbkbJpb7"`
	BintangDbkbJpb7 *string  `json:"bintangDbkbJpb7"`
	LantaiMinJpb7   *int     `json:"lantaiMinJpb7"`
	LantaiMaxJpb7   *int     `json:"lantaiMaxJpb7"`
	NilaiDbkbJpb7   *float64 `json:"nilaiDbkbJpb7"`
}

type UpdateDto struct {
	Provinsi_Kode   *string  `json:"provinsi_kode"`
	Daerah_Kode     *string  `json:"daerah_kode"`
	TahunDbkbJpb7   *string  `json:"tahunDbkbJpb7"`
	JenisDbkbJpb7   *string  `json:"jenisDbkbJpb7"`
	BintangDbkbJpb7 *string  `json:"bintangDbkbJpb7"`
	LantaiMinJpb7   *int     `json:"lantaiMinJpb7"`
	LantaiMaxJpb7   *int     `json:"lantaiMaxJpb7"`
	NilaiDbkbJpb7   *float64 `json:"nilaiDbkbJpb7"`
}

type FilterDto struct {
	Provinsi_Kode   *string  `json:"provinsi_kode"`
	Daerah_Kode     *string  `json:"daerah_kode"`
	TahunDbkbJpb7   *string  `json:"tahunDbkbJpb7"`
	JenisDbkbJpb7   *string  `json:"jenisDbkbJpb7"`
	BintangDbkbJpb7 *string  `json:"bintangDbkbJpb7"`
	LantaiMinJpb7   *int     `json:"lantaiMinJpb7"`
	LantaiMaxJpb7   *int     `json:"lantaiMaxJpb7"`
	NilaiDbkbJpb7   *float64 `json:"nilaiDbkbJpb7"`
	Page            int      `json:"page"`
	PageSize        int      `json:"page_size"`
}
