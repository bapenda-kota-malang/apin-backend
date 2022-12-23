package firebird

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

// create new connection to firebird database
//
// example:
//
//	connString := "user:password@servername/foo/bar.fdb"
//	fdb, err := firebird.NewConn(connString)
func NewConn(connString string) (conn *sql.DB, err error) {
	conn, err = sql.Open("firebirdsql", connString)
	if err != nil {
		return nil, fmt.Errorf("new firebird conn: %w", err)
	}
	return
}
