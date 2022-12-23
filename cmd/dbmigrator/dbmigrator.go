package main

import (
	"log"

	"github.com/bapenda-kota-malang/apin-backend/pkg/db/firebird"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db/oracle"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db/postgres"
)

func main() {
	postgreConn := ""
	pdb, err := postgres.NewConn(postgreConn)
	if err != nil {
		log.Fatal(err)
	}
	defer pdb.Close()

	firebirdConn := ""
	fdb, err := firebird.NewConn(firebirdConn)
	if err != nil {
		log.Fatal(err)
	}
	defer fdb.Close()

	oracleDsn := ""
	odb, err := oracle.NewConn(oracleDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer odb.Close()
}
