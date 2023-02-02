/*
This flow
*/
package main

import (
	"log"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/seeder"
)

func main() {
	// load config
	var c = &seeder.Config{}
	if err := c.LoadConfig(); err != nil {
		log.Fatalf("load config: %s", err)
	}

	opt := seeder.NewOption()

	if *opt.UpdateSql {
		if err := seeder.UpdateSqlFile(); err != nil {
			log.Fatalf("update sql file: %s", err)
		}
	}

	if err := seeder.RunImportSql(c.GenerateCommand(opt.Notrx)); err != nil {
		log.Fatalf("update sql file: %s", err)
	}

	log.Println("seeding done")
}
