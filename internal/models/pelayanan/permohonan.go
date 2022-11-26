package permohonan

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

var (
	JenisPelayanan = [...]string{"0001", "0002", "0003", "0004"}
)

type PstPermohonan struct {
	Id                     uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	NoPelayanan            *string         `json:"noPelayanan" gorm:"type:varchar(50)"`
	NOP                    *string         `json:"nop" gorm:"type:varchar(50)"`
	KanwilId               *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId                *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan         *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan        *string         `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan        *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	NoSuratPermohonan      *string         `json:"noSuratPermohonan" gorm:"type:varchar(50)"`
	TanggalSuratPermohonan *datatypes.Date `json:"tanggalSuratPermohonan"`
	NamaPemohon            *string         `json:"namaWP" gorm:"type:varchar(30)"`
	AlamatPemohon          *string         `json:"letakOP" gorm:"type:varchar(40)"`
	Keterangan             *string         `json:"keterangan" gorm:"type:varchar(75)"`
	CatatanPermohonan      *string         `json:"catatan" gorm:"type:varchar(75)"`
	StatusKolektif         *string         `json:"statusKolektif" gorm:"type:varchar(1)"`
	TanggalTerima          *datatypes.Date `json:"tanggalTerima"`
	NIP                    *string         `json:"nip" gorm:"type:varchar(9)"`
	PenerimaanBerkas       *string         `json:"penerimaanBerkas" gorm:"type:varchar(50)"`
	gormhelper.DateModel

	PstDataOPBaru            *PstDataOPBaru            `json:"nopBaru" gorm:"foreignKey:Id;references:PermohonanId"`
	PstDetail                *PstDetail                `json:"nopDetail" gorm:"foreignKey:Id;references:PermohonanId"`
	PstPermohonanPengurangan *PstPermohonanPengurangan `json:"nopPengurangan" gorm:"foreignKey:Id;references:PermohonanId"`
}

type PermohonanRequestDto struct {
	KanwilId              *string `json:"kanwilId"`
	KppbbId               *string `json:"kppbbId"`
	NoPelayanan           *string `json:"noPelayanan"`
	StatusKolektif        *string `json:"statusKolektif"`
	NoSuratPermohonan     *string `json:"noSuratPermohonan"`
	JenisPelayanan        *string `json:"jenisPelayanan"`
	TanggalTerima         *string `json:"tanggalTerima"`
	TanggalSelesai        *string `json:"tanggalSelesai"`
	TanggalPermohonan     *string `json:"tanggalPermohonan"`
	NOP                   *string `json:"nop"`
	NamaWP                *string `json:"namaWP"`
	LetakOP               *string `json:"letakOP"`
	Keterangan            *string `json:"keterangan"`
	TahunPajak            *string `json:"tahunPajak"`
	PenerimaanBerkas      *string `json:"penerimaanBerkas"`
	Catatan               *string `json:"catatan"`
	NIP                   *string `json:"nip"`
	JenisPengurangan      *string `json:"jenisPengurangan"`
	PersentasePengurangan *string `json:"persentasePengurangan"`
	SeksiBerkasID         *string `json:"seksiBerkasID"`
	CatatanPenyerahan     *string `json:"catatanPenyerahan"`
	NamaPenerima          *string `json:"namaPenerima"`
	NIPPenyerah           *string `json:"nipPenyerah"`
	StatusSelesai         *int    `json:"statusSelesai"`
	TanggalPenyerahan     *string `json:"tanggalPenyerahan"`
}

type FilterDto struct {
	StatusKolektif    *string         `json:"statusKolektif"`
	NoSuratPermohonan *string         `json:"noSuratPermohonan"`
	JenisPelayanan    *string         `json:"jenisPelayanan"`
	TanggalTerima     *datatypes.Date `json:"tanggalTerima"`
	TanggalSelesai    *datatypes.Date `json:"tanggalSelesai"`
	TanggalPermohonan *datatypes.Date `json:"tanggalPermohonan"`
	NOP               *string         `json:"nop"`
	NamaWP            *string         `json:"namaWP"`
	LetakOP           *string         `json:"letakOP"`
	TahunPajak        *string         `json:"tahunPajak"`
	PenerimaanBerkas  *string         `json:"penerimaanBerkas"`
}

type PstDetail struct {
	Id                    uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string         `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string         `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string         `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string         `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string         `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string         `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string         `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string         `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	JenisPelayananID      *string         `json:"jenisPelayananID" gorm:"type:varchar(4)"`
	TahunPajakPemohon     *string         `json:"tahunPajakPemohon" gorm:"type:varchar(30)"`
	NIP                   *string         `json:"nip" gorm:"type:varchar(9)"`
	Notes                 *string         `json:"catatan" gorm:"type:varchar(75)"`
	StatusSelesai         *int            `json:"keterangan" gorm:"type:integer"`
	TanggalSelesai        *datatypes.Date `json:"tanggalSelesai"`
	SeksiBerkasID         *string         `json:"seksiBerkasID" gorm:"type:varchar(2)"`
	TanggalPenyerahan     *datatypes.Date `json:"tanggalPenyerahan"`
	gormhelper.DateModel
}

type PstDataOPBaru struct {
	Id                    uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64 `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	NamaWPBaru            *string `json:"namaWP" gorm:"type:varchar(30)"`
	LetakOPBaru           *string `json:"letakOP" gorm:"type:varchar(40)"`
	gormhelper.DateModel
}

type PstPermohonanPengurangan struct {
	Id                    uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64  `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string  `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string  `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string  `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string  `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string  `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string  `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string  `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string  `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string  `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string  `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string  `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string  `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	JenisPengurangan      *string  `json:"jenisPengurangan" gorm:"type:varchar(1)"`
	PersentasePengurangan *float64 `json:"persentasePengurangan" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type permohonanNOP struct {
	PermohonanProvinsiID  string
	PermohonanKotaID      string
	PermohonanKecamatanID string
	PermohonanKelurahanID string
	PermohonanBlokID      string
	NoUrutPemohon         string
	PemohonJenisOPID      string
}

type PermohonanNOPResponse struct {
	NOP                   *string `json:"nop"`
	NamaWP                *string `json:"namaWP"`
	LetakOP               *string `json:"letakOP"`
	Keterangan            *string `json:"keterangan"`
	TahunPajak            *string `json:"tahunPajak"`
	TanggalSelesai        *string `json:"tanggalSelesai"`
	JenisPengurangan      *string `json:"jenisPengurangan"`
	PersentasePengurangan *string `json:"persentasePengurangan"`
}

func DecodeNOPPermohonan(nop *string) *permohonanNOP {
	if nop != nil {
		var tempNOP string
		tempNOP = *nop
		result := permohonanNOP{
			PermohonanProvinsiID:  tempNOP[0:2],
			PermohonanKotaID:      tempNOP[2:4],
			PermohonanKecamatanID: tempNOP[4:7],
			PermohonanKelurahanID: tempNOP[7:10],
			PermohonanBlokID:      tempNOP[10:13],
			NoUrutPemohon:         tempNOP[13:17],
			PemohonJenisOPID:      tempNOP[17:18],
		}
		return &result
	}
	return nil
}

func (i PstPermohonan) GetDataPermohonanNoPelayanan() string {
	result := fmt.Sprintf("%s%s%s", *i.TahunPelayanan, *i.BundelPelayanan, *i.NoUrutPelayanan)
	return result
}

func (i PermohonanRequestDto) SetDataPermohonanRequestDtoTransformer(noUrut *string) PstPermohonan {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	if noUrut == nil {
		if i.NoPelayanan != nil {
			tempNoUrut := *i.NoPelayanan
			tempNoUrut = tempNoUrut[9:]
			noUrut = &tempNoUrut
		}
	}
	result := PstPermohonan{
		// NoPelayanan:            i.NoPelayanan,
		NOP:                    i.NOP,
		KanwilId:               i.KanwilId,
		KppbbId:                i.KppbbId,
		TahunPelayanan:         &year,
		BundelPelayanan:        i.JenisPelayanan,
		NoUrutPelayanan:        noUrut,
		NoSuratPermohonan:      i.NoSuratPermohonan,
		TanggalSuratPermohonan: (*datatypes.Date)(gormhelper.StringToDate(*i.TanggalPermohonan)),
		NamaPemohon:            i.NamaWP,
		AlamatPemohon:          i.LetakOP,
		Keterangan:             i.Keterangan,
		CatatanPermohonan:      i.Catatan,
		StatusKolektif:         i.StatusKolektif,
		TanggalTerima:          (*datatypes.Date)(gormhelper.StringToDate(*i.TanggalTerima)),
		NIP:                    i.NIP,
		PenerimaanBerkas:       i.PenerimaanBerkas,
	}
	return result
}

func (i PstPermohonan) GetDataPermohonanRequestDtoTransformer(nop *PermohonanNOPResponse) PermohonanRequestDto {
	var (
		tanggalTerimaTemp     *string
		tanggalPermohonanTemp *string
	)
	tempNoPelayanan := i.GetDataPermohonanNoPelayanan()

	temp := (*time.Time)(i.TanggalTerima).String()
	tanggalTerimaTemp = &temp
	temp = (*time.Time)(i.TanggalSuratPermohonan).String()
	tanggalPermohonanTemp = &temp

	result := PermohonanRequestDto{
		KanwilId:          i.KanwilId,
		KppbbId:           i.KppbbId,
		NoPelayanan:       &tempNoPelayanan,
		StatusKolektif:    i.StatusKolektif,
		NoSuratPermohonan: i.NoSuratPermohonan,
		JenisPelayanan:    i.BundelPelayanan,
		TanggalTerima:     tanggalTerimaTemp,
		TanggalPermohonan: tanggalPermohonanTemp,
		PenerimaanBerkas:  i.PenerimaanBerkas,
		NIP:               i.NIP,
		Catatan:           i.CatatanPermohonan,
		Keterangan:        i.Keterangan,
	}
	if nop != nil {
		result.NOP = nop.NOP
		result.TanggalSelesai = nop.TanggalSelesai
		result.NamaWP = nop.NamaWP
		result.LetakOP = nop.LetakOP
		if nop.Keterangan != nil {
			result.Keterangan = nop.Keterangan
		}
		result.TahunPajak = nop.TahunPajak
		result.JenisPengurangan = nop.JenisPengurangan
		result.PersentasePengurangan = nop.PersentasePengurangan
	}
	return result
}

func (i PstPermohonan) SetDataPermohonanTransformer(req PermohonanRequestDto) (*PstDataOPBaru, *PstDetail, *PstPermohonanPengurangan) {
	tempNOP := DecodeNOPPermohonan(req.NOP)
	if *i.BundelPelayanan == JenisPelayanan[0] {
		data := PstDataOPBaru{
			PermohonanId:          &i.Id,
			KanwilId:              i.KanwilId,
			KppbbId:               i.KanwilId,
			TahunPelayanan:        i.TahunPelayanan,
			BundelPelayanan:       i.BundelPelayanan,
			NoUrutPelayanan:       i.NoUrutPelayanan,
			PermohonanProvinsiID:  &tempNOP.PermohonanProvinsiID,
			PermohonanKotaID:      &tempNOP.PermohonanKotaID,
			PermohonanKecamatanID: &tempNOP.PermohonanKecamatanID,
			PermohonanKelurahanID: &tempNOP.PermohonanKelurahanID,
			PermohonanBlokID:      &tempNOP.PermohonanBlokID,
			NoUrutPemohon:         &tempNOP.NoUrutPemohon,
			PemohonJenisOPID:      &tempNOP.PemohonJenisOPID,
			NamaWPBaru:            i.NamaPemohon,
			LetakOPBaru:           i.AlamatPemohon,
		}
		return &data, nil, nil
	} else if *i.BundelPelayanan == JenisPelayanan[1] {
		// currentTime := time.Now()
		// datenow, _ := time.Parse("2006-01-02", currentTime.Format("2006-01-02"))
		data := PstDetail{
			PermohonanId:          &i.Id,
			KanwilId:              i.KanwilId,
			KppbbId:               i.KanwilId,
			TahunPelayanan:        i.TahunPelayanan,
			BundelPelayanan:       i.BundelPelayanan,
			NoUrutPelayanan:       i.NoUrutPelayanan,
			PermohonanProvinsiID:  &tempNOP.PermohonanProvinsiID,
			PermohonanKotaID:      &tempNOP.PermohonanKotaID,
			PermohonanKecamatanID: &tempNOP.PermohonanKecamatanID,
			PermohonanKelurahanID: &tempNOP.PermohonanKelurahanID,
			PermohonanBlokID:      &tempNOP.PermohonanBlokID,
			NoUrutPemohon:         &tempNOP.NoUrutPemohon,
			PemohonJenisOPID:      &tempNOP.PemohonJenisOPID,
			JenisPelayananID:      i.BundelPelayanan,
			TahunPajakPemohon:     req.TahunPajak,
			NIP:                   req.NIPPenyerah,
			Notes:                 req.CatatanPenyerahan,
			StatusSelesai:         req.StatusSelesai,
			TanggalSelesai:        (*datatypes.Date)(gormhelper.StringToDate(*req.TanggalSelesai)),
			SeksiBerkasID:         req.SeksiBerkasID,
			TanggalPenyerahan:     (*datatypes.Date)(gormhelper.StringToDate(*req.TanggalPenyerahan)),
		}
		return nil, &data, nil
	} else if *i.BundelPelayanan == JenisPelayanan[2] {
		tempPengurangan, _ := strconv.ParseFloat(strings.TrimSpace(*req.PersentasePengurangan), 64)
		data := PstPermohonanPengurangan{
			PermohonanId:          &i.Id,
			KanwilId:              i.KanwilId,
			KppbbId:               i.KanwilId,
			TahunPelayanan:        i.TahunPelayanan,
			BundelPelayanan:       i.BundelPelayanan,
			NoUrutPelayanan:       i.NoUrutPelayanan,
			PermohonanProvinsiID:  &tempNOP.PermohonanProvinsiID,
			PermohonanKotaID:      &tempNOP.PermohonanKotaID,
			PermohonanKecamatanID: &tempNOP.PermohonanKecamatanID,
			PermohonanKelurahanID: &tempNOP.PermohonanKelurahanID,
			PermohonanBlokID:      &tempNOP.PermohonanBlokID,
			NoUrutPemohon:         &tempNOP.NoUrutPemohon,
			PemohonJenisOPID:      &tempNOP.PemohonJenisOPID,
			JenisPengurangan:      req.JenisPengurangan,
			PersentasePengurangan: &tempPengurangan,
		}
		return nil, nil, &data
	}
	return nil, nil, nil
}

func (i PstDataOPBaru) GetDataPermohonanNOP() string {
	result := fmt.Sprintf("%s%s%s%s%s%s%s", *i.PermohonanProvinsiID, *i.PermohonanKotaID, *i.PermohonanKecamatanID, *i.PermohonanKelurahanID, *i.PermohonanBlokID, *i.NoUrutPemohon, *i.PemohonJenisOPID)
	return result
}

func (i PstDetail) GetDataPermohonanNOP() string {
	result := fmt.Sprintf("%s%s%s%s%s%s%s", *i.PermohonanProvinsiID, *i.PermohonanKotaID, *i.PermohonanKecamatanID, *i.PermohonanKelurahanID, *i.PermohonanBlokID, *i.NoUrutPemohon, *i.PemohonJenisOPID)
	return result
}

func (i PstPermohonanPengurangan) GetDataPermohonanNOP() string {
	result := fmt.Sprintf("%s%s%s%s%s%s%s", *i.PermohonanProvinsiID, *i.PermohonanKotaID, *i.PermohonanKecamatanID, *i.PermohonanKelurahanID, *i.PermohonanBlokID, *i.NoUrutPemohon, *i.PemohonJenisOPID)
	return result
}

func (i PstDataOPBaru) GetNOPResponse() *PermohonanNOPResponse {
	nop := i.GetDataPermohonanNOP()
	result := PermohonanNOPResponse{
		NOP:        &nop,
		NamaWP:     i.NamaWPBaru,
		LetakOP:    i.LetakOPBaru,
		TahunPajak: i.TahunPelayanan,
	}
	return &result
}

func (i PstDetail) GetNOPResponse() *PermohonanNOPResponse {
	temp := (*time.Time)(i.TanggalSelesai).String()
	tanggalSelesaiTemp := &temp

	nop := i.GetDataPermohonanNOP()
	result := PermohonanNOPResponse{
		NOP:            &nop,
		TahunPajak:     i.TahunPajakPemohon,
		Keterangan:     i.Notes,
		TanggalSelesai: tanggalSelesaiTemp,
	}
	return &result
}

func (i PstPermohonanPengurangan) GetNOPResponse() *PermohonanNOPResponse {
	nop := i.GetDataPermohonanNOP()
	persentase := fmt.Sprint(i.PersentasePengurangan)
	result := PermohonanNOPResponse{
		NOP:                   &nop,
		TahunPajak:            i.TahunPelayanan,
		JenisPengurangan:      i.JenisPengurangan,
		PersentasePengurangan: &persentase,
	}
	return &result
}
