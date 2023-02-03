package seeder

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/seeder"
)

type seedSqlHandler struct {
	SeedSqlService seeder.SeedSqlService
}

func NewSeedSqlHandler(seedSqlService seeder.SeedSqlService) *seedSqlHandler {
	return &seedSqlHandler{SeedSqlService: seedSqlService}
}

func (s *seedSqlHandler) UpdateSqlFile() error {
	// get list files
	files, err := s.SeedSqlService.GetListFile()
	if err != nil {
		return fmt.Errorf("get list file: %w", err)
	}

	// create seed.sql for make one file sql
	if err := s.SeedSqlService.WriteSeedSql(files); err != nil {
		return fmt.Errorf("write sql file: %w", err)
	}
	return nil
}

func (s *seedSqlHandler) ImportSql(cmd string) error {
	return s.SeedSqlService.RunImportSql(cmd)
}
