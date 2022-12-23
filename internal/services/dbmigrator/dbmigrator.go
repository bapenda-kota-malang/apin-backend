package dbmigrator

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbmigrator"
	"github.com/jackc/pgx/v4"
)

// make data string to double quotes
//
// ex: hello, world! -> "hello, world!"
func quoteString(data string) string {
	n := len(data)
	b := make([]byte, n+2)
	b[0] = '"'
	b[len(b)-1] = '"'
	for i := 0; i < n; i++ {
		b[i+1] = data[i]
	}
	return string(b)
}

type DbMigratorService struct {
	postgre    *pgx.Conn
	firebird   *sql.DB
	oracle     *sql.DB
	dbMigrator []dbmigrator.Migrator
}

func NewMigratorService(postgre *pgx.Conn, firebird *sql.DB, oracle *sql.DB) dbmigrator.DbMigratorService {
	return &DbMigratorService{postgre: postgre, firebird: firebird, oracle: oracle}
}

// copy from src table to dst table
//
// note: before run this must run ParseJson(jsonPath)
func (s *DbMigratorService) RunCopyData() error {
	var (
		srcRows *sql.Rows
		err     error
	)
	if len(s.dbMigrator) == 0 {
		return fmt.Errorf("db migrator data empty")
	}
	for _, v := range s.dbMigrator {
		if v.Dst.Dialect != "postgresql" {
			return fmt.Errorf("unknown dst dialect: %s for table: %s", v.Dst.Dialect, v.Dst.Table.Name)
		}
		if len(v.Src.Table.Column) != len(v.Dst.Table.Column) {
			return fmt.Errorf("length column not same, src: %v, dst: %v", len(v.Src.Table.Column), len(v.Dst.Table.Column))
		}

		// get data source
		src := v.Src
		srcQuery := fmt.Sprintf("SELECT %s FROM %s", strings.Join(src.Table.Column, ", "), src.Table.Name)
		switch src.Dialect {
		case "firebird":
			srcRows, err = s.firebird.QueryContext(context.Background(), srcQuery)
		case "oracle":
			srcRows, err = s.oracle.QueryContext(context.Background(), srcQuery)
		default:
			log.Printf("unknown src dialect: %s for table: %s", src.Dialect, src.Table.Name)
			continue
		}
		if err != nil {
			log.Printf("error table: %s, src query: %s", src.Table.Name, err)
			continue
		}

		// process data from source to destination, use copy protocol to bulk insert
		dst := v.Dst
		copyCount, err := s.postgre.CopyFrom(
			context.Background(),
			pgx.Identifier{dst.Table.Name},
			dst.Table.Column,
			&copyBeetweenTable{columns: src.Table.Column, srcRows: srcRows},
		)
		if err != nil {
			log.Printf("error table: %s, dst copy: %s", src.Table.Name, err)
			continue
		}
		log.Printf("success copy from: %s to: %s with %v datas", src.Table.Name, dst.Table.Name, copyCount)
	}
	return nil
}

// parse json bytes to db migrator struct
func (s *DbMigratorService) ParseJson(b []byte) error {
	var data []dbmigrator.Migrator
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	s.dbMigrator = data
	return nil
}
