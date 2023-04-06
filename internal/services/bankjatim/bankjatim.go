package bankjatim

import (
	"context"

	mlbj "github.com/bapenda-kota-malang/apin-backend/internal/models/logbankjatim"
	slbj "github.com/bapenda-kota-malang/apin-backend/internal/services/logbankjatim"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
)

const source = "bank jatim"

func Registration(ctx context.Context, payload bankjatim.RequestRegistration, userId uint64) error {
	var resp interface{}
	var err error
	respCode := ""
	vaResp, err := a.BankJatim.Va.Registration(ctx, payload)
	resp = err
	if err == nil {
		resp = vaResp
		respCode = string(vaResp.Status.ResponseCode)
	}

	// save to log db bank jatim
	dataLog := mlbj.CreateDynamicDto{
		Path:         "registration",
		Content:      payload,
		Response:     resp,
		ResponseCode: respCode,
		User_Id:      userId,
	}
	if _, err := slbj.Create(dataLog); err != nil {
		return err
	}
	return err
}

func Posting(ctx context.Context, req bankjatim.RequestPosting) (*bankjatim.ResponsePosting, error) {
	var data bankjatim.ResponsePosting
	return &data, nil
}
