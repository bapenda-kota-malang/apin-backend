package seeder

import (
	"fmt"
	"os"
	"path"

	"github.com/bapenda-kota-malang/apin-backend/internal/services/seeder"
)

func UpdateSqlFile() error {
	// get path
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get wd path: %w", err)
	}

	// get list files

	files, err := seeder.GetListFile(path.Join(wd, "sqls"))
	if err != nil {
		return fmt.Errorf("get list file: %w", err)
	}

	// create seed.sql for make one file sql
	if err := seeder.WriteSeedSql(files); err != nil {
		return fmt.Errorf("write sql file: %w", err)
	}
	return nil
}

func RunImportSql(cmd string) error {
	if err := seeder.ImportToDb(cmd); err != nil {
		return err
	}
	return nil
}
