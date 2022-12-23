package firebird

import (
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func NewConn(connString string) (conn *sql.DB, err error) {
	conn, err = sql.Open("firebirdsql", connString)
	if err != nil {
		return nil, fmt.Errorf("new firebird conn: %w", err)
	}
	return
}
