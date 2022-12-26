package firebird

import (
	"database/sql"

	dbsql "github.com/bapenda-kota-malang/apin-backend/pkg/db/sql"
	_ "github.com/nakagami/firebirdsql"
)

// create new connection to firebird database
//
// example:
//
//	connString := "user:password@servername/foo/bar.fdb"
//	fdb, err := firebird.NewConn(connString)
func NewConn(connString string) (conn *sql.DB, err error) {
	return dbsql.NewGenericCon("firebirdsql", connString)
}
