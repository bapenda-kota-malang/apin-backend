/*
This flow
*/
package main

import (
	"log"
	"os"

	"github.com/bapenda-kota-malang/apin-backend/internal/handlers/seeder"
	_seederHandler "github.com/bapenda-kota-malang/apin-backend/internal/handlers/seeder"
	_seederService "github.com/bapenda-kota-malang/apin-backend/internal/services/seeder"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db"
)

// func initDb(c *seeder.Config) (*gorm.DB, error) {
// 	dbConf := db.NewDbConf("postgres")
// 	if err := dbConf.GenerateDsn(c.Host, c.DbUser, c.Password, c.DbName, c.Port); err != nil {
// 		return nil, err
// 		// log.Fatalf("generate dsn db: %s", err)
// 	}
// 	db, err := dbConf.InitDb()
// 	if err != nil {
// 		return nil, err
// 		// log.Fatalf("init db: %s", err)
// 	}
// 	return db, nil
// }

func main() {
	// load config
	var c = seeder.NewConfig()
	if err := c.LoadConfig(); err != nil {
		log.Fatalf("load config: %s", err)
	}

	// load option
	opt := seeder.NewOption()
	opt.ParseOption()

	// get path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("get wd path: %s", err)
	}

	sSqlService := _seederService.NewSeed(wd)
	sSqlHandler := _seederHandler.NewSeedSqlHandler(sSqlService)

	// update seed.sql list data from inside sqls folder
	if *opt.UpdateSql {
		if err := sSqlHandler.UpdateSqlFile(); err != nil {
			log.Fatalf("update sql file: %s", err)
		}
	}

	// update seed.sql list store procedure
	if *opt.StoreProcedure == "sync" {
		dbConf := db.NewDbConf("", "postgres")
		dbConf.GenerateDsn(c.Host, c.DbUser, c.Password, c.DbName, c.Port)
		db, err := dbConf.InitDb()
		if err != nil {
			log.Fatalf("setup db err %s", err)
		}
		sSpService := _seederService.NewSeedSp(db)
		_, err = sSpService.GetListFunction()
		if err != nil {
			log.Fatalf("get list sp from db err %s", err)
		}
	}

	if err := sSqlHandler.ImportSql(c.GenerateCommand(opt.Notrx)); err != nil {
		log.Fatalf("import sql file: %s", err)
	}

	log.Println("seeding done")
}
