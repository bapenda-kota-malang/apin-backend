package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type PstLampiran struct {
	Id                      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan" gorm:"type:varchar(2)"`
	PermohonanId            *uint64   `json:"permohonanId" gorm:"type:integer"`
	KanwilId                *string   `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId                 *string   `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan          *string   `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan         *string   `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan" gorm:"TEXT"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa" gorm:"TEXT"`
	LampiranSK              *string   `json:"lampiranSK" gorm:"TEXT"`
	LampiranKTP             *string   `json:"lampiranKTP" gorm:"TEXT"`
	LampiranKK              *string   `json:"lampiranKK" gorm:"TEXT"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah" gorm:"TEXT"`
	LampiranSppt            *string   `json:"lampiranSppt" gorm:"TEXT"`
	LampiranImb             *string   `json:"lampiranImb" gorm:"TEXT"`
	LampiranAkte            *string   `json:"lampiranAkte" gorm:"TEXT"`
	LampiranSkPensiun       *string   `json:"lampiranSkPensiun" gorm:"TEXT"`
	LampiranSpptStts        *string   `json:"lampiranSpptStts" gorm:"TEXT"`
	LampiranSTTS            *string   `json:"lampiranSTTS" gorm:"TEXT"`
	LampiranSkPengurangan   *string   `json:"lampiranSkPengurangan" gorm:"TEXT"`
	LampiranSkKeberatan     *string   `json:"lampiranSkKeberatan" gorm:"TEXT"`
	LampiranSkkpPbb         *string   `json:"lampiranSkkpPbb" gorm:"TEXT"`
	LampiranSpmkpPbb        *string   `json:"lampiranSpmkpPbb" gorm:"TEXT"`
	LampiranLainLain        *string   `json:"lampiranLainLain" gorm:"TEXT"`
	LampiranLaporanKeuangan *string   `json:"lampiranLaporanKeuangan" gorm:"TEXT"`
	LampiranLikuid          *string   `json:"lampiranLikuid" gorm:"TEXT"`
	LampiranSlipGaji        *string   `json:"lampiranSlipGaji" gorm:"TEXT"`
	LampiranListrik         *string   `json:"lampiranListrik" gorm:"TEXT"`
	LampiranLetakOP         *string   `json:"lampiranLetakOP" gorm:"TEXT"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik" gorm:"TEXT"`
	LampiranFotoOP          *string   `json:"lampiranFotoOP" gorm:"TEXT"`
	PermohonanNOP
	gormhelper.DateModel
}

type PstLampiranCreateDTO struct {
	Id                      uuid.UUID `json:"id"`
	PermohonanId            *uint64   `json:"permohonanId"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan"`
	KanwilId                *string   `json:"kanwilId"`
	KppbbId                 *string   `json:"kppbbId"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
	BundelPelayanan         *string   `json:"bundlePelayanan"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa"`
	LampiranSK              *string   `json:"lampiranSK"`
	LampiranKTP             *string   `json:"lampiranKTP"`
	LampiranKK              *string   `json:"lampiranKK"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah"`
	LampiranSppt            *[]string `json:"lampiranSppt"`
	LampiranImb             *string   `json:"lampiranImb"`
	LampiranAkte            *string   `json:"lampiranAkte"`
	LampiranSkPensiun       *string   `json:"lampiranSkPensiun"`
	LampiranSpptStts        *string   `json:"lampiranSpptStts"`
	LampiranSTTS            *string   `json:"lampiranSTTS"`
	LampiranSkPengurangan   *string   `json:"lampiranSkPengurangan"`
	LampiranSkKeberatan     *string   `json:"lampiranSkKeberatan"`
	LampiranSkkpPbb         *string   `json:"lampiranSkkpPbb"`
	LampiranSpmkpPbb        *string   `json:"lampiranSpmkpPbb"`
	LampiranLainLain        *string   `json:"lampiranLainLain"`
	LampiranLaporanKeuangan *string   `json:"lampiranLaporanKeuangan"`
	LampiranLikuid          *string   `json:"lampiranLikuid"`
	LampiranSlipGaji        *string   `json:"lampiranSlipGaji"`
	LampiranListrik         *string   `json:"lampiranListrik"`
	LampiranLetakOP         *string   `json:"lampiranLetakOP"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik"`
	LampiranFotoOP          *[]string `json:"lampiranFotoOP"`
	PermohonanNOP
	gormhelper.DateModel
}

type PstLampiranUpdateDTO struct {
	Id                      uuid.UUID `json:"id"`
	PermohonanId            *uint64   `json:"permohonanId"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan"`
	KanwilId                *string   `json:"kanwilId"`
	KppbbId                 *string   `json:"kppbbId"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
	BundelPelayanan         *string   `json:"bundlePelayanan"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa"`
	LampiranSK              *string   `json:"lampiranSK"`
	LampiranKTP             *string   `json:"lampiranKTP"`
	LampiranKK              *string   `json:"lampiranKK"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah"`
	LampiranSppt            *[]string `json:"lampiranSppt"`
	LampiranImb             *string   `json:"lampiranImb"`
	LampiranAkte            *string   `json:"lampiranAkte"`
	LampiranSkPensiun       *string   `json:"lampiranSkPensiun"`
	LampiranSpptStts        *string   `json:"lampiranSpptStts"`
	LampiranSTTS            *string   `json:"lampiranSTTS"`
	LampiranSkPengurangan   *string   `json:"lampiranSkPengurangan"`
	LampiranSkKeberatan     *string   `json:"lampiranSkKeberatan"`
	LampiranSkkpPbb         *string   `json:"lampiranSkkpPbb"`
	LampiranSpmkpPbb        *string   `json:"lampiranSpmkpPbb"`
	LampiranLainLain        *string   `json:"lampiranLainLain"`
	LampiranLaporanKeuangan *string   `json:"lampiranLaporanKeuangan"`
	LampiranLikuid          *string   `json:"lampiranLikuid"`
	LampiranSlipGaji        *string   `json:"lampiranSlipGaji"`
	LampiranListrik         *string   `json:"lampiranListrik"`
	LampiranLetakOP         *string   `json:"lampiranLetakOP"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik"`
	LampiranFotoOP          *[]string `json:"lampiranFotoOP"`
	PermohonanNOP
	gormhelper.DateModel
}
