package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb3 struct {
	Id               uint64  `json:"id" gorm:"primarykey;autoIncrement"`
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	gh.DateModel
}

type RegJpb3CreateDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb3UpdateDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb3FilterDto struct {
	PstPermohonan_id *uint64 `json:"pstPermohonan_id"`
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

func (input *RegOpbJpb3CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb3CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb3CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb3CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb3CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb3CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
