package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type RegPstLampiran struct {
	Id                      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PermohonanId            *uint64   `json:"permohonanId" gorm:"type:integer"`
	KanwilId                *string   `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId                 *string   `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan          *string   `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan         *string   `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan" gorm:"type:varchar(255)"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa" gorm:"type:varchar(255)"`
	LampiranKTP             *string   `json:"lampiranKTP" gorm:"type:varchar(255)"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah" gorm:"type:varchar(255)"`
	LampiranSppt            *string   `json:"lampiranSppt" gorm:"type:varchar(255)"`
	LampiranImb             *string   `json:"lampiranImb" gorm:"type:varchar(255)"`
	LampiranAkte            *string   `json:"lampiranAkte" gorm:"type:varchar(255)"`
	LampiranSkPensiun       *string   `json:"lampiranSkPensiun" gorm:"type:varchar(255)"`
	LampiranSpptStts        *string   `json:"lampiranSpptStts" gorm:"type:varchar(255)"`
	LampiranSTTS            *string   `json:"lampiranSTTS" gorm:"type:varchar(255)"`
	LampiranSkPengurangan   *string   `json:"lampiranSkPengurangan" gorm:"type:varchar(255)"`
	LampiranSkKeberatan     *string   `json:"lampiranSkKeberatan" gorm:"type:varchar(255)"`
	LampiranSkkpPbb         *string   `json:"lampiranSkkpPbb" gorm:"type:varchar(255)"`
	LampiranSpmkpPbb        *string   `json:"lampiranSpmkpPbb" gorm:"type:varchar(255)"`
	LampiranLainLain        *string   `json:"lampiranLainLain" gorm:"type:varchar(255)"`
	LampiranLikuid          *string   `json:"lampiranLikuid" gorm:"type:varchar(255)"`
	LampiranListrik         *string   `json:"lampiranListrik" gorm:"type:varchar(255)"`
	PermohonanNOP
	gormhelper.DateModel
}

type RegPstLampiranCreateDTO struct {
	Id                      uuid.UUID `json:"id"`
	PermohonanId            *uint64   `json:"permohonanId"`
	KanwilId                *string   `json:"kanwilId"`
	KppbbId                 *string   `json:"kppbbId"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
	BundelPelayanan         *string   `json:"bundlePelayanan"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa"`
	LampiranKTP             *string   `json:"lampiranKTP"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah"`
	LampiranSppt            *string   `json:"lampiranSppt"`
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
	LampiranLikuid          *string   `json:"lampiranLikuid"`
	LampiranListrik         *string   `json:"lampiranListrik"`
	PermohonanNOP
	gormhelper.DateModel
}

type RegPstLampiranUpdateDTO struct {
	Id                      uuid.UUID `json:"id"`
	PermohonanId            *uint64   `json:"permohonanId"`
	KanwilId                *string   `json:"kanwilId"`
	KppbbId                 *string   `json:"kppbbId"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
	BundelPelayanan         *string   `json:"bundlePelayanan"`
	NoUrutPelayanan         *string   `json:"noUrutPelayanan"`
	LampiranPermohonan      *string   `json:"lampiranPermohonan"`
	LampiranSuratKuasa      *string   `json:"lampiranSuratKuasa"`
	LampiranKTP             *string   `json:"lampiranKTP"`
	LampiranSertifikatTanah *string   `json:"lampiranSertifikatTanah"`
	LampiranSppt            *string   `json:"lampiranSppt"`
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
	LampiranLikuid          *string   `json:"lampiranLikuid"`
	LampiranListrik         *string   `json:"lampiranListrik"`
	PermohonanNOP
	gormhelper.DateModel
}
