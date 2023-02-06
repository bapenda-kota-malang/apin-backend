package seeder

import (
	"fmt"
	"runtime"

	pkgConfig "github.com/bapenda-kota-malang/apin-backend/pkg/config"
)

// config struct specific for seeder, extend load config method from pkg
type config struct {
	Host     string
	Port     string
	DbUser   string
	Password string
	DbName   string
}

func NewConfig() *config {
	return &config{}
}

func (c *config) LoadConfig() error {
	return pkgConfig.LoadConfig(c)
}

func (c *config) GenerateCommand(notrx *bool) string {
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DbUser, c.Password, c.Host, c.Port, c.DbName)
	trxCommand := " -1"
	if *notrx {
		trxCommand = ""
	}
	cmd := ""

	switch runtime.GOOS {
	case "windows":
		cmd = fmt.Sprintf("psql%s -f seed.sql %s", trxCommand, uri)
	case "linux":
		cmd = fmt.Sprintf("psql %s%s -f seed.sql", uri, trxCommand)
	}
	return cmd
}
