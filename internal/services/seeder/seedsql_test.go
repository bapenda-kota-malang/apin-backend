package seeder

import (
	"reflect"
	"testing"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/seeder"
	"github.com/bapenda-kota-malang/apin-backend/pkg/db"
	"gorm.io/gorm"
)

func Test_seedSqlService_getListFunction(t *testing.T) {
	dbConf := db.NewDbConf("host=127.0.0.1 port=5432 user=dexwip password=sultan81199 dbname=apin sslmode=disable TimeZone=Asia/Jakarta", "postgres")
	db, err := dbConf.InitDb()
	if err != nil {
		t.Errorf("setup err %s", err)
	}
	type fields struct {
		basePath string
		Db       *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []seeder.StoreProcedure
		wantErr bool
	}{
		{name: "test", fields: fields{Db: db}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &seedSqlService{
				basePath: tt.fields.basePath,
				Db:       tt.fields.Db,
			}
			got, err := s.getListFunction()
			if (err != nil) != tt.wantErr {
				t.Errorf("seedSqlService.getListFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("seedSqlService.getListFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
