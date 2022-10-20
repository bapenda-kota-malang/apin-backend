package indukobjekpajak

import (
	opp "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type IndukObjekPajak struct {
	opp.DetailLspop
	gh.DateModel
}
