package harga

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type HargaSatuan struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Propinsi_Id    string  `json:"propinsi_Id"`
	Dati2_Id       string  `json:"dati2_Id"`
	TahunSatuan    string  `json:"tahunSatuan"`
	Pekerjaan_Kode string  `json:"pekerjaan_kode"`
	Kegiatan_Kode  string  `json:"kegiatan_kode"`
	HargaSatuan    float64 `json:"hargaSatuan"`
	gh.DateModel
}

type HargaResource struct {
	Id                 uint64  `json:"id" gorm:"primaryKey"`
	Propinsi_Id        string  `json:"propinsi_Id"`
	Dati2_Id           string  `json:"dati2_Id"`
	TahunResource      string  `json:"tahunResource"`
	GroupResource_Kode string  `json:"groupResource_kode"`
	Resource_Kode      string  `json:"resource_kode"`
	Kanwil_Id          *string `json:"kanwil_Id"`
	KPPBB_Id           *string `json:"kppbb_Id"`
	JenisDokumen       string  `json:"jenisDokumen"`
	NoDokumen          string  `json:"noDokumen"`
	HargaResource      float64 `json:"hargaResource"`
	gh.DateModel
}

type HargaKegiatan struct {
	Id                  uint64  `json:"id" gorm:"primaryKey"`
	Propinsi_Id         string  `json:"propinsi_Id"`
	Dati2_Id            string  `json:"dati2_Id"`
	TahunKegiatan       string  `json:"tahunKegiatan"`
	JPB_Kode            string  `json:"jpb_kode"`
	Bangunan_Tipe       string  `json:"bangunan_tipe"`
	BangunanLantai_Kode string  `json:"bangunanLantai_kode"`
	Pekerjaan_Kode      string  `json:"pekerjaan_kode"`
	Kegiatan_Kode       string  `json:"kegiatan_kode"`
	HargaKegiatan       float64 `json:"hargaKegiatan"`
	gh.DateModel
}

type HargaKegiatanJPB8 struct {
	Id                uint64  `json:"id" gorm:"primaryKey"`
	Propinsi_Id       string  `json:"propinsi_Id"`
	Dati2_Id          string  `json:"dati2_Id"`
	TahunKegiatanJPB8 string  `json:"tahunKegiatanJPB8"`
	Pekerjaan_Kode    string  `json:"pekerjaan_kode"`
	Kegiatan_Kode     string  `json:"kegiatan_kode"`
	HRGLBRBENTMIN     float64 `json:"HRGLBRBENTMIN"`
	HRGLBRBENTMAX     float64 `json:"HRGLBRBENTMAX"`
	HRGTINGKOLOMMIN   float64 `json:"HRGTINGKOLOMMIN"`
	HRGTINGKOLOMMAX   float64 `json:"HRGTINGKOLOMMAX"`
	HargaKegiatanJPB8 float64 `json:"hargaKegiatanJPB8"`
	gh.DateModel
}
