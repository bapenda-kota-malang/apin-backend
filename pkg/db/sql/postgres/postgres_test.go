package postgres

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestNewConn(t *testing.T) {
	type args struct {
		connString string
	}
	tests := []struct {
		name          string
		args          args
		wantConn      *pgx.Conn
		wantErr       bool
		errConclusion string
	}{
		{name: "error invalid dsn", args: args{connString: "postgre://username:password@localhost:5432/database_name"}, wantConn: nil, wantErr: true, errConclusion: "invalid dsn"},
		{name: "error connection refused", args: args{connString: "postgres://username:password@localhost:22222/database_name"}, wantConn: nil, wantErr: true, errConclusion: "connection refused"},
		{name: "error unknown hostname", args: args{connString: "postgres://username:password@unknown.host.test:22222/database_name"}, wantConn: nil, wantErr: true, errConclusion: "no such host"},
		{name: "error password failed", args: args{connString: "postgres://aa:password@localhost:5432/database_name"}, wantConn: nil, wantErr: true, errConclusion: "SQLSTATE 28P01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConn, err := NewConn(tt.args.connString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && !strings.Contains(err.Error(), tt.errConclusion) {
				t.Errorf("NewConn() error = %v, errConclusion %v", err, tt.errConclusion)
				return
			}
			if !reflect.DeepEqual(gotConn, tt.wantConn) {
				t.Errorf("NewConn() = %v, want %v", gotConn, tt.wantConn)
				return
			}
		})
	}
}
