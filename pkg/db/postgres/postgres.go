package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// create new connection to postgresql database
//
// example:
//
//	connString := "postgres://username:password@localhost:5432/database_name"
//	pdb, err := postgres.NewConn(connString)
func NewConn(connString string) (conn *sql.DB, err error) {
	conn, err = sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("new postgre conn: %w", err)
	}
	return
}
