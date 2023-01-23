package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sc "github.com/jinzhu/copier"
)

type RegJpb8 struct {
	Id               uint64  `json:"id" gorm:"primarykey;autoIncrement"`
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	gh.DateModel
}

type RegJpb8CreateDto struct {
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb8UpdateDto struct {
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb8FilterDto struct {
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

func (input *RegOpbJpb8CreateDto) GetFasilitasBangunan() *mrfb.CreateDto {
	return input.RegFasilitasBangunans
}

func (input *RegOpbJpb8CreateDto) GetNop() *string {
	return input.Nop
}

func (input *RegOpbJpb8CreateDto) GetTanggalPendataan() *string {
	return input.TanggalPendataan
}

func (input *RegOpbJpb8CreateDto) GetTanggalPemeriksaan() *string {
	return input.TanggalPemeriksaan
}

func (input *RegOpbJpb8CreateDto) GetTanggalPerekaman() *string {
	return input.TanggalPerekaman
}

func (input *RegOpbJpb8CreateDto) GetObjekPajakBangunan() (CreateDto, error) {
	var data CreateDto
	if err := sc.Copy(&data, &input); err != nil {
		return CreateDto{}, err
	}
	return data, nil
}
