package oracle

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// create new connection to oracle database
//
// example:
//
//	dataSourceName := `user="scott" password="tiger" connectString="dbhost:1521/orclpdb1"`
//	odb, err := oracle.NewConn(dataSourceName)
//
// The connectString can be ANYTHING that SQL*Plus or Oracle Call Interface (OCI) accepts:
//
// a service name an https://download.oracle.com/ocomdocs/global/Oracle-Net-19c-Easy-Connect-Plus.pdf like "host:port/service_name"
//
// or a connect descriptor like (DESCRIPTION=...).
//
// You can specify connection timeout seconds with "?connect_timeout=15" - Ping uses this timeout, NOT the Deadline in Context! Note that connect_timeout requires at least 19c client.
//
// For more connection options, see https://godror.github.io/godror/doc/connection.html.
func NewConn(dataSourceName string) (conn *sql.DB, err error) {
	conn, err = sql.Open("godror", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("new oracle conn: %w", err)
	}
	return
}
