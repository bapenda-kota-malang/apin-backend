package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewConn(connString string) (conn *sql.DB, err error) {
	conn, err = sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("new postgre conn: %w", err)
	}
	return
}
