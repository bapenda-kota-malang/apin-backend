package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type RegPstLampiran struct {
	Id                      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan" gorm:"type:varchar(2)"`
	PermohonanId            *uint64   `json:"permohonanId" gorm:"type:integer"`
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
	LampiranLikuid          *string   `json:"lampiranLikuid" gorm:"TEXT"`
	LampiranListrik         *string   `json:"lampiranListrik" gorm:"TEXT"`
	LampiranLetakOP         *string   `json:"lampiranLetakOP" gorm:"TEXT"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik" gorm:"TEXT"`
	LampiranFotoOP          *string   `json:"lampiranFotoOP" gorm:"TEXT"`
	PermohonanNOP
	gormhelper.DateModel
}

type RegPstLampiranCreateDTO struct {
	PermohonanId            *uint64   `json:"permohonanId"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
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
	LampiranLikuid          *string   `json:"lampiranLikuid"`
	LampiranListrik         *string   `json:"lampiranListrik"`
	LampiranLetakOP         *string   `json:"lampiranLetakOP"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik"`
	LampiranFotoOP          *[]string `json:"lampiranFotoOP"`
	PermohonanNOP
	gormhelper.DateModel
}

type RegPstLampiranUpdateDTO struct {
	Id                      uuid.UUID `json:"id"`
	PermohonanId            *uint64   `json:"permohonanId"`
	BuktiKepemilikan        *string   `json:"buktiKepemilikan"`
	TahunPelayanan          *string   `json:"tahunPelayanan"`
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
	LampiranLetakOP         *string   `json:"lampiranLetakOP"`
	LampiranHakMilik        *string   `json:"lampiranHakMilik"`
	LampiranFotoOP          *[]string `json:"lampiranFotoOP"`
	PermohonanNOP
	gormhelper.DateModel
}
