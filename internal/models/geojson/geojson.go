package geojson

import (
	"strconv"
	"strings"

	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type GeoJson struct {
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	FID_        *string `json:"FID_" gorm:"type:varchar(20)"`
	Entity      *string `json:"Entity" gorm:"type:varchar(20)"`
	Layer       *string `json:"Layer" gorm:"type:varchar(20)"`
	Color       *string `json:"Color" gorm:"type:varchar(20)"`
	Linetype    *string `json:"Linetype" gorm:"type:varchar(20)"`
	Elevation   *string `json:"Elevation" gorm:"type:varchar(20)"`
	LineWt      *string `json:"LineWt" gorm:"type:varchar(20)"`
	RefName     *string `json:"RefName" gorm:"type:varchar(20)"`
	D_Nop       *string `json:"d_nop" gorm:"type:varchar(20)"`
	D_Luas      *string `json:"d_luas" gorm:"type:varchar(20)"`
	Nama        *string `json:"NAMA" gorm:"type:varchar(20)"`
	No          *string `json:"NO" gorm:"type:varchar(20)"`
	Alamat      *string `json:"ALAMAT" gorm:"type:varchar(20)"`
	Rt          *string `json:"RT" gorm:"type:varchar(20)"`
	Rw          *string `json:"RW" gorm:"type:varchar(20)"`
	Kelurahan   *string `json:"KELURAHAN" gorm:"type:varchar(20)"`
	Kecamatan   *string `json:"KECAMATAN" gorm:"type:varchar(20)"`
	Kota        *string `json:"KOTA" gorm:"type:varchar(20)"`
	Luas_tanah  *string `json:"LUAS_TANAH" gorm:"type:varchar(20)"`
	Kode_ZNT    *string `json:"KODE_ZNT" gorm:"type:varchar(20)"`
	Jnis_Tanah  *string `json:"JNIS_TANAH" gorm:"type:varchar(20)"`
	J_Bangunan  *string `json:"J_BANGUNAN" gorm:"type:varchar(20)"`
	L_Bangun    *string `json:"L_BANGUN" gorm:"type:varchar(20)"`
	Njop_Bumi   *string `json:"NJOP_BUMI" gorm:"type:varchar(20)"`
	Njop_Bgn    *string `json:"NJOP_BGN" gorm:"type:varchar(20)"`
	Njop_Total  *string `json:"NJOP_TOTAL" gorm:"type:varchar(20)"`
	Pendataan   *string `json:"PENDATAAN" gorm:"type:varchar(20)"`
	Keterangan  *string `json:"KETERANGAN" gorm:"type:varchar(20)"`
	Edit_Kav    *string `json:"EDIT_KAV" gorm:"type:varchar(20)"`
	Fungsi      *string `json:"FUNGSI" gorm:"type:varchar(20)"`
	Jnis_Pajak  *string `json:"JNIS_PAJAK" gorm:"type:varchar(20)"`
	Coordinates *string `json:"Coordinates" gorm:"type:varchar(250)"`
	gh.DateModel
}

type CreateDto struct {
	FID_        *string `json:"FID_"`
	Entity      *string `json:"Entity"`
	Layer       *string `json:"Layer"`
	Color       *string `json:"Color"`
	Linetype    *string `json:"Linetype"`
	Elevation   *string `json:"Elevation"`
	LineWt      *string `json:"LineWt"`
	RefName     *string `json:"RefName"`
	D_Nop       *string `json:"d_nop"`
	D_Luas      *string `json:"d_luas"`
	Nama        *string `json:"NAMA"`
	No          *string `json:"NO"`
	Alamat      *string `json:"ALAMAT"`
	Rt          *string `json:"RT"`
	Rw          *string `json:"RW"`
	Kelurahan   *string `json:"KELURAHAN"`
	Kecamatan   *string `json:"KECAMATAN"`
	Kota        *string `json:"KOTA"`
	Luas_tanah  *string `json:"LUAS_TANAH"`
	Kode_ZNT    *string `json:"KODE_ZNT"`
	Jnis_Tanah  *string `json:"JNIS_TANAH"`
	J_Bangunan  *string `json:"J_BANGUNAN"`
	L_Bangun    *string `json:"L_BANGUN"`
	Njop_Bumi   *string `json:"NJOP_BUMI"`
	Njop_Bgn    *string `json:"NJOP_BGN"`
	Njop_Total  *string `json:"NJOP_TOTAL"`
	Pendataan   *string `json:"PENDATAAN"`
	Keterangan  *string `json:"KETERANGAN"`
	Edit_Kav    *string `json:"EDIT_KAV"`
	Fungsi      *string `json:"FUNGSI"`
	Jnis_Pajak  *string `json:"JNIS_PAJAK"`
	Coordinates *string `json:"Coordinates"`
}

type UpdateDto struct {
	FID_        *string `json:"FID_"`
	Entity      *string `json:"Entity"`
	Layer       *string `json:"Layer"`
	Color       *string `json:"Color"`
	Linetype    *string `json:"Linetype"`
	Elevation   *string `json:"Elevation"`
	LineWt      *string `json:"LineWt"`
	RefName     *string `json:"RefName"`
	D_Nop       *string `json:"d_nop"`
	D_Luas      *string `json:"d_luas"`
	Nama        *string `json:"NAMA"`
	No          *string `json:"NO"`
	Alamat      *string `json:"ALAMAT"`
	Rt          *string `json:"RT"`
	Rw          *string `json:"RW"`
	Kelurahan   *string `json:"KELURAHAN"`
	Kecamatan   *string `json:"KECAMATAN"`
	Kota        *string `json:"KOTA"`
	Luas_tanah  *string `json:"LUAS_TANAH"`
	Kode_ZNT    *string `json:"KODE_ZNT"`
	Jnis_Tanah  *string `json:"JNIS_TANAH"`
	J_Bangunan  *string `json:"J_BANGUNAN"`
	L_Bangun    *string `json:"L_BANGUN"`
	Njop_Bumi   *string `json:"NJOP_BUMI"`
	Njop_Bgn    *string `json:"NJOP_BGN"`
	Njop_Total  *string `json:"NJOP_TOTAL"`
	Pendataan   *string `json:"PENDATAAN"`
	Keterangan  *string `json:"KETERANGAN"`
	Edit_Kav    *string `json:"EDIT_KAV"`
	Fungsi      *string `json:"FUNGSI"`
	Jnis_Pajak  *string `json:"JNIS_PAJAK"`
	Coordinates *string `json:"Coordinates"`
}

type FilterDto struct {
	// dynamic fields
	FID_       *string `json:"FID_"`
	Entity     *string `json:"Entity"`
	Layer      *string `json:"Layer"`
	Color      *string `json:"Color"`
	Linetype   *string `json:"Linetype"`
	Elevation  *string `json:"Elevation"`
	LineWt     *string `json:"LineWt"`
	RefName    *string `json:"RefName"`
	D_Nop      *string `json:"d_nop"`
	D_Luas     *string `json:"d_luas"`
	Nama       *string `json:"NAMA"`
	No         *string `json:"NO"`
	Alamat     *string `json:"ALAMAT"`
	Rt         *string `json:"RT"`
	Rw         *string `json:"RW"`
	Kelurahan  *string `json:"KELURAHAN"`
	Kecamatan  *string `json:"KECAMATAN"`
	Kota       *string `json:"KOTA"`
	Luas_tanah *string `json:"LUAS_TANAH"`
	Kode_ZNT   *string `json:"KODE_ZNT"`
	Jnis_Tanah *string `json:"JNIS_TANAH"`
	J_Bangunan *string `json:"J_BANGUNAN"`
	L_Bangun   *string `json:"L_BANGUN"`
	Njop_Bumi  *string `json:"NJOP_BUMI"`
	Njop_Bgn   *string `json:"NJOP_BGN"`
	Njop_Total *string `json:"NJOP_TOTAL"`
	Pendataan  *string `json:"PENDATAAN"`
	Keterangan *string `json:"KETERANGAN"`
	Edit_Kav   *string `json:"EDIT_KAV"`
	Fungsi     *string `json:"FUNGSI"`
	Jnis_Pajak *string `json:"JNIS_PAJAK"`
	// fixed fields
	Page         int   `json:"page"`
	PageSize     int64 `json:"page_size"`
	NoPagination bool  `json:"no_pagination"`
}

func batchActions(a []float64, c int) [][]float64 {
	r := (len(a) + c - 1) / c
	b := make([][]float64, r)
	lo, hi := 0, c
	for i := range b {
		if hi > len(a) {
			hi = len(a)
		}
		b[i] = a[lo:hi:hi]
		lo, hi = hi, hi+c
	}
	return b
}

func (i GeoJson) GeoJsonPosition() [][][]float64 {
	tempC := i.Coordinates
	tempArrayString := strings.Split(*tempC, ",")
	var (
		tempArrayFloat []float64
		tempVal        float64
	)

	for i := range tempArrayString {
		tempVal, _ = strconv.ParseFloat(tempArrayString[i], 64)
		tempArrayFloat = append(tempArrayFloat, tempVal)
	}
	tempResult := batchActions(tempArrayFloat, 3)
	result := [][][]float64{tempResult}
	return result
}
