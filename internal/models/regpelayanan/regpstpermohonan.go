package regpstpermohonan

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"

	ori "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	mroppbb "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
)

var (
	JenisPelayanan = [...]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}
)

type RegPstPermohonan struct {
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
	Status                 *string         `json:"status"`
	User_ID                *uint64         `json:"user_ID"`
	gormhelper.DateModel
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
	AlasanPengurangan     *string `json:"alasanPengurangan"`
	PersentasePengurangan *string `json:"persentasePengurangan"`
	SeksiBerkasID         *string `json:"seksiBerkas"`
	CatatanPenyerahan     *string `json:"catatanPenyerahan"`
	NamaPenerima          *string `json:"namaPenerima"`
	NIPPenyerah           *string `json:"nipPenyerah"`
	StatusSelesai         *int    `json:"statusSelesai"`
	Status                *string `json:"status"`
	User_ID               *uint64 `json:"user_ID"`
	TanggalPenyerahan     *string `json:"tanggalPenyerahan"`

	PstOpjekPajakPBB *mroppbb.CreateDto `json:"oppbb"`
}

type PermohonanApprovalRequestDto struct {
	Id                    *uint64 `json:"id"`
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
	AlasanPengurangan     *string `json:"alasanPengurangan"`
	PersentasePengurangan *string `json:"persentasePengurangan"`
	SeksiBerkasID         *string `json:"seksiBerkas"`
	CatatanPenyerahan     *string `json:"catatanPenyerahan"`
	NamaPenerima          *string `json:"namaPenerima"`
	NIPPenyerah           *string `json:"nipPenyerah"`
	StatusSelesai         *int    `json:"statusSelesai"`
	Status                *string `json:"status"`
	User_ID               *uint64 `json:"user_ID"`
	TanggalPenyerahan     *string `json:"tanggalPenyerahan"`
	CatatanApproval       *string `json:"catatanApproval"`
	Jabatan               *string `json:"jabatan"`
	Divisi                *string `json:"divisi"`

	PstDataOPBaru            *RegPstDataOPBaru            `json:"pstBaru"`
	PstDetail                *RegPstDetail                `json:"pstDetil"`
	PstPermohonanPengurangan *RegPstPermohonanPengurangan `json:"pstPengurangan"`
	PstLampiran              *RegPstLampiran              `json:"pstLampiran"`
	PstLogApproval           *[]ori.PstLogApproval        `json:"pstLogApproval"`

	PstOpjekPajakPBB *mroppbb.CreateDto `json:"oppbb"`
}

type FilterDto struct {
	StatusKolektif    *string         `json:"statusKolektif"`
	NoSuratPermohonan *string         `json:"noSuratPermohonan"`
	BundelPelayanan   *[]string       `json:"jenisPelayanan"`
	TanggalTerima     *datatypes.Date `json:"tanggalTerima"`
	TanggalSelesai    *datatypes.Date `json:"tanggalSelesai"`
	TanggalPermohonan *datatypes.Date `json:"tanggalPermohonan"`
	NOP               *string         `json:"nop"`
	NamaWP            *string         `json:"namaWP" refsource:"NamaPemohon"`
	NamaWP_Opt        *string         `json:"namaWP_opt"`
	LetakOP           *string         `json:"letakOP"`
	TahunPajak        *string         `json:"tahunPajak"`
	PenerimaanBerkas  *string         `json:"penerimaanBerkas"`
	Page              int             `json:"page"`
	PageSize          int             `json:"page_size"`
}

type PermohonanNOP struct {
	PermohonanProvinsiID  string `json:"provinsiKode" gorm:"type:varchar(2)"`
	PermohonanKotaID      string `json:"daerahKode" gorm:"type:varchar(2)"`
	PermohonanKecamatanID string `json:"kecamatanKode" gorm:"type:varchar(3)"`
	PermohonanKelurahanID string `json:"kelurahanKode" gorm:"type:varchar(3)"`
	PermohonanBlokID      string `json:"blokKode" gorm:"type:varchar(3)"`
	NoUrutPemohon         string `json:"noUrutKode" gorm:"type:varchar(4)"`
	PemohonJenisOPID      string `json:"jenisOP" gorm:"type:varchar(1)"`
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

type PstPermohonanResponse struct {
	Id                     uint64          `json:"id"`
	NoPelayanan            *string         `json:"noPelayanan"`
	NOP                    *string         `json:"nop"`
	KanwilId               *string         `json:"kanwilId"`
	KppbbId                *string         `json:"kppbbId"`
	TahunPelayanan         *string         `json:"tahunPelayanan"`
	BundelPelayanan        *string         `json:"bundlePelayanan"`
	NoUrutPelayanan        *string         `json:"noUrutPelayanan"`
	NoSuratPermohonan      *string         `json:"noSuratPermohonan"`
	TanggalSuratPermohonan *datatypes.Date `json:"tanggalSuratPermohonan"`
	NamaPemohon            *string         `json:"namaWP"`
	AlamatPemohon          *string         `json:"letakOP"`
	Keterangan             *string         `json:"keterangan"`
	CatatanPermohonan      *string         `json:"catatan"`
	StatusKolektif         *string         `json:"statusKolektif"`
	TanggalTerima          *datatypes.Date `json:"tanggalTerima"`
	NIP                    *string         `json:"nip"`
	PenerimaanBerkas       *string         `json:"penerimaanBerkas"`
	Status                 *string         `json:"status"`
	CatatanApproval        *string         `json:"catatanApproval"`

	PstDataOPBaru            *RegPstDataOPBaru            `json:"pstBaru"`
	PstDetail                *RegPstDetail                `json:"pstDetil"`
	PstPermohonanPengurangan *RegPstPermohonanPengurangan `json:"pstPengurangan"`
	PstLampiran              *RegPstLampiran              `json:"pstLampiran"`
	PstLogApproval           *[]ori.PstLogApproval        `json:"pstLogApproval"`
	PstLogApprovalRes        *ori.ResponsePSTLogApproval  `json:"pstLogApprovalRes"`

	// PembetulanSpptSKPSTP     *PembetulanSpptSKPSTP     `json:"pembetulanSpptSKPSTP"`
	// PembatalanSppt           *PembatalanSppt           `json:"pembatalanSppt"`
	// KeputusanKeberatanPbb    *KeputusanKeberatanPbb    `json:"keputusanKeberatanPbb"`
	// SPMKP                    *SPMKP                    `json:"spmkp"`
	// SkSk                     *SkSk                     `json:"sksk"`

	PstOpjekPajakPBB *mroppbb.CreateDto `json:"oppbb"`
}

type RequestApprovalPermohonan struct {
	Id          uint64  `json:"id"`
	NoPelayanan *string `json:"noPelayanan"`
	NOP         *string `json:"nop"`
	NIP         *string `json:"nip"`
	Status      *string `json:"status"`
	User_ID     *uint64 `json:"user_id"`
	Catatan     *string `json:"catatan"`
	Keterangan  *string `json:"keterangan"`
	Jabatan     *string `json:"jabatan"`
	Divisi      *string `json:"divisi"`
	gormhelper.DateModel
}

func DecodeNOPPermohonan(nop *string) *PermohonanNOP {
	if nop != nil {
		var tempNOP string
		tempNOP = *nop
		result := PermohonanNOP{
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

func (i RegPstPermohonan) GetDataPermohonanNoPelayanan() string {
	result := fmt.Sprintf("%s%s%s", *i.TahunPelayanan, *i.BundelPelayanan, *i.NoUrutPelayanan)
	return result
}

func (i RegPstPermohonan) SetPstPermohonanResponse() PstPermohonanResponse {
	result := PstPermohonanResponse{
		Id:                     i.Id,
		NoPelayanan:            i.NoPelayanan,
		NOP:                    i.NOP,
		KanwilId:               i.KanwilId,
		KppbbId:                i.KppbbId,
		TahunPelayanan:         i.TahunPelayanan,
		BundelPelayanan:        i.BundelPelayanan,
		NoUrutPelayanan:        i.NoUrutPelayanan,
		NoSuratPermohonan:      i.NoSuratPermohonan,
		TanggalSuratPermohonan: i.TanggalSuratPermohonan,
		NamaPemohon:            i.NamaPemohon,
		AlamatPemohon:          i.AlamatPemohon,
		Keterangan:             i.Keterangan,
		CatatanPermohonan:      i.CatatanPermohonan,
		StatusKolektif:         i.StatusKolektif,
		TanggalTerima:          i.TanggalTerima,
		NIP:                    i.NIP,
		PenerimaanBerkas:       i.PenerimaanBerkas,
	}
	return result
}

func (i PermohonanRequestDto) SetDataPermohonanRequestDtoTransformer(noUrut *string) RegPstPermohonan {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	if noUrut == nil {
		if i.NoPelayanan != nil {
			tempNoUrut := *i.NoPelayanan
			tempNoUrut = tempNoUrut[9:]
			noUrut = &tempNoUrut
		}
	}
	var tempTglTerima *datatypes.Date
	if i.TanggalTerima == nil {
		tempTglTerima = nil
	} else {
		tempTglTerima = (*datatypes.Date)(gormhelper.StringToDate(*i.TanggalTerima))
	}
	result := RegPstPermohonan{
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
		TanggalTerima:          tempTglTerima,
		NIP:                    i.NIP,
		PenerimaanBerkas:       i.PenerimaanBerkas,
	}
	return result
}

func (i RegPstPermohonan) GetDataPermohonanRequestDtoTransformer(nop *PermohonanNOPResponse) PermohonanRequestDto {
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

func (i RegPstPermohonan) SetDataPermohonanTransformer(req PermohonanRequestDto) (*RegPstDataOPBaru, *RegPstDetailInput, *RegPstPermohonanPengurangan, *PermohonanNOP) {
	var (
		// tempTglPenyerahan *datatypes.Date
		tempTglSelesai *datatypes.Date
	)

	tempNOP := DecodeNOPPermohonan(req.NOP)
	// if req.TanggalPenyerahan != nil {
	// 	tempTglPenyerahan = (*datatypes.Date)(gormhelper.StringToDate(*req.TanggalPenyerahan))
	// }
	if req.TanggalSelesai != nil {
		tempTglSelesai = (*datatypes.Date)(gormhelper.StringToDate(*req.TanggalSelesai))
	}

	detail := RegPstDetailInput{
		PermohonanId:      &i.Id,
		KanwilId:          i.KanwilId,
		KppbbId:           i.KanwilId,
		TahunPelayanan:    i.TahunPelayanan,
		BundelPelayanan:   i.BundelPelayanan,
		NoUrutPelayanan:   i.NoUrutPelayanan,
		JenisPelayananID:  i.BundelPelayanan,
		TahunPajakPemohon: req.TahunPajak,
		TanggalSelesai:    tempTglSelesai,
		// NIP:                   req.NIPPenyerah,
		// Notes:                 req.CatatanPenyerahan,
		// StatusSelesai:         req.StatusSelesai,
		// SeksiBerkasID:         req.SeksiBerkasID,
		// TanggalPenyerahan:     tempTglPenyerahan,
	}

	if *i.BundelPelayanan == JenisPelayanan[0] {
		data := RegPstDataOPBaru{
			PermohonanId:    &i.Id,
			KanwilId:        i.KanwilId,
			KppbbId:         i.KanwilId,
			TahunPelayanan:  i.TahunPelayanan,
			BundelPelayanan: i.BundelPelayanan,
			NoUrutPelayanan: i.NoUrutPelayanan,
			NamaWPBaru:      i.NamaPemohon,
			LetakOPBaru:     i.AlamatPemohon,
		}
		return &data, &detail, nil, tempNOP
	} else if *i.BundelPelayanan == JenisPelayanan[7] || *i.BundelPelayanan == JenisPelayanan[9] {
		tempPengurangan, _ := strconv.ParseFloat(strings.TrimSpace(*req.PersentasePengurangan), 64)
		data := RegPstPermohonanPengurangan{
			PermohonanId:          &i.Id,
			KanwilId:              i.KanwilId,
			KppbbId:               i.KanwilId,
			TahunPelayanan:        i.TahunPelayanan,
			BundelPelayanan:       i.BundelPelayanan,
			NoUrutPelayanan:       i.NoUrutPelayanan,
			JenisPengurangan:      req.JenisPengurangan,
			AlasanPengurangan:     req.AlasanPengurangan,
			PersentasePengurangan: &tempPengurangan,
		}
		return nil, &detail, &data, tempNOP
	}
	return nil, &detail, nil, tempNOP
}
