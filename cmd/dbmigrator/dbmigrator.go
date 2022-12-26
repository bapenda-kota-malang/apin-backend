package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sdbm "github.com/bapenda-kota-malang/apin-backend/internal/services/dbmigrator"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db/sql/firebird"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db/sql/oracle"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db/sql/postgres"
	"github.com/spf13/viper"
)

type DbConf struct {
	Postgres string
	Firebird string
	Oracle   string
}

// load from yml files to struct config
func (c *DbConf) loadConfig() (err error) {
	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// read the file
	if err = viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("error reading config file: %s", err)
		return
	}

	// map to app
	if err = viper.Unmarshal(&c); err != nil {
		err = fmt.Errorf("unable to decode into struct: %s", err)
		return
	}
	return
}

func main() {
	// load config
	var c = &DbConf{}
	if err := c.loadConfig(); err != nil {
		log.Fatalf("load config: %s", err)
	}

	pdb, err := postgres.NewConn(c.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer pdb.Close(context.Background())

	fdb, err := firebird.NewConn(c.Firebird)
	if err != nil {
		log.Fatal(err)
	}
	defer fdb.Close()

	odb, err := oracle.NewConn(c.Oracle)
	if err != nil {
		log.Fatal(err)
	}
	defer odb.Close()

	// create service and execute
	dbMigratorService := sdbm.NewMigratorService(pdb, fdb, odb)
	b, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := dbMigratorService.ParseJson(b); err != nil {
		log.Fatal(err)
	}
	if err := dbMigratorService.RunCopyData(); err != nil {
		log.Fatal(err)
	}
	log.Println("done migrate from source db to destination db")
}
