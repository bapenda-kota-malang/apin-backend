package objekpajakbangunan

import (
	mfb "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type Jpb3 struct {
	nop.NopDetail
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	gh.DateModel
}

type Jpb3CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type Jpb3UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type Jpb3FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

func (input *OpbJpb3CreateDto) GetFasilitasBangunan() *mfb.CreateDto {
	return input.FasilitasBangunans
}

func (input *OpbJpb3CreateDto) GetNop() *string {
	return input.Nop
}

func (input *OpbJpb3CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *OpbJpb3CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *OpbJpb3CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *OpbJpb3CreateDto) GetObjekPajakBangunan() (*CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return nil, err
	}
	return &data, nil
}
