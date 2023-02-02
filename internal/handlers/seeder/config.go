package seeder

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/pkg/config"
)

// config struct specific for seeder, extend load config method from pkg
type Config struct {
	Host     string
	Port     string
	DbUser   string
	Password string
	DbName   string
}

func (c *Config) LoadConfig() error {
	return config.LoadConfig(c)
}

func (c *Config) GenerateCommand(notrx *bool) string {
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.DbUser, c.Password, c.Host, c.Port, c.DbName)
	trxCommand := " -1"
	if *notrx {
		trxCommand = ""
	}
	return fmt.Sprintf("psql %s%s -f seed.sql", uri, trxCommand)
}
