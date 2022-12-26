package sql

import (
	"database/sql"
	"fmt"
)

func NewGenericCon(driverName, dataSourceName string) (*sql.DB, error) {
	conn, err := sql.Open(driverName, dataSourceName)
	if err == nil {
		err = conn.Ping()
	}

	if err != nil {
		return nil, fmt.Errorf("new %s conn: %w", driverName, err)
	}
	return conn, nil
}
