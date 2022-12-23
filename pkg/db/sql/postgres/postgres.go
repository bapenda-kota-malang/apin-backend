package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// create new connection to postgresql database using pgx connection
//
// example:
//
//	connString := "postgres://username:password@localhost:5432/database_name"
//	pdb, err := postgres.NewConn(connString)
func NewConn(connString string) (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("new pgx conn: %w", err)
	}
	return
}
