package pengurangan

import (
	mpst "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	spstdetail "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan/pstdetail"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

func checkPstDetail(tahun, bundel, noUrut string) ([]mpst.PstDetail, error) {
	respPstDetail, err := spstdetail.GetByNoPelayanan(mpst.GetByNoPelayananPstDetailDto{
		TahunPelayanan:  tahun,
		BundelPelayanan: bundel,
		NoUrutPelayanan: noUrut,
	})
	if err != nil || (respPstDetail == nil && err == nil) {
		return nil, err
	}
	return respPstDetail.(rp.OKSimple).Data.([]mpst.PstDetail), nil
}
