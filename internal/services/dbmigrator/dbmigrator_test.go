package dbmigrator

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbmigrator"
	"github.com/jackc/pgx/v4"
)

func Test_quoteString(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "success quote string", args: args{data: "test"}, want: `"test"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quoteString(tt.args.data); got != tt.want {
				t.Errorf("quouteString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMigratorService(t *testing.T) {
	type args struct {
		postgre  *pgx.Conn
		firebird *sql.DB
		oracle   *sql.DB
	}
	tests := []struct {
		name string
		args args
		want dbmigrator.DbMigratorService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMigratorService(tt.args.postgre, tt.args.firebird, tt.args.oracle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMigratorService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbMigratorService_RunCopyData(t *testing.T) {
	type fields struct {
		postgre    *pgx.Conn
		firebird   *sql.DB
		oracle     *sql.DB
		dbMigrator []dbmigrator.Migrator
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DbMigratorService{
				postgre:    tt.fields.postgre,
				firebird:   tt.fields.firebird,
				oracle:     tt.fields.oracle,
				dbMigrator: tt.fields.dbMigrator,
			}
			if err := s.RunCopyData(); (err != nil) != tt.wantErr {
				t.Errorf("DbMigratorService.CopyData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDbMigratorService_ParseJson(t *testing.T) {
	jsonEmpty := `[]`
	jsonNotArray := `{}`
	jsonDataNotObject := `["data", "failed"]`
	type args struct {
		b []byte
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantErrMessage string
		wantDbMigrator []dbmigrator.Migrator
	}{
		{name: "error end of json", args: args{b: []byte("")}, wantErr: true, wantErrMessage: "unexpected end of JSON input"},
		{name: "success but json empty", args: args{b: []byte(jsonEmpty)}, wantErr: false, wantDbMigrator: []dbmigrator.Migrator{}},
		{name: "error parse json not array", args: args{b: []byte(jsonNotArray)}, wantErr: true, wantErrMessage: "cannot unmarshal"},
		{name: "error parse json data not object", args: args{b: []byte(jsonDataNotObject)}, wantErr: true, wantErrMessage: "cannot unmarshal"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DbMigratorService{}
			err := s.ParseJson(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DbMigratorService.ParseJson() error = %v, wantErr %v, wantErrMessage %v", err, tt.wantErr, tt.wantErrMessage)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.wantErrMessage) {
				t.Errorf("DbMigratorService.ParseJson() error = %v, wantErrMessage %v", err, tt.wantErrMessage)
				return
			}
			if !reflect.DeepEqual(s.dbMigrator, tt.wantDbMigrator) {
				t.Errorf("s.dbMigrator = %v, want %v", s.dbMigrator, tt.wantDbMigrator)
			}
		})
	}
}
