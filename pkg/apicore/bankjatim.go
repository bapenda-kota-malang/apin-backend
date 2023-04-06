package apicore

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/integration"
	"github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
	"go.uber.org/zap"
)

type bankJatimConf struct {
	Url string
}

type BankJatimClient struct {
	Va *bankjatim.VirtualAccount
}

func (a *app) initBankJatimClient() {
	if a.BankJatimConf.Url == "" {
		Logger.Info("instantiation",
			zap.String("type", "integration"),
			zap.String("source", "bank jatim"),
			zap.String("status", "unset"),
			zap.String("message", "no URL provided"))
		return
	}
	httpClient := integration.NewClient()
	BankJatim = &BankJatimClient{
		Va: bankjatim.NewVirtualAccount(a.BankJatimConf.Url, httpClient),
	}
}
