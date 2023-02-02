package seeder

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/seeder"
	"gorm.io/gorm"
)

type storeProcedureService struct {
	Db           *gorm.DB
	listFunction *[]seeder.StoreProcedure
}

func NewstoreProcedureService(db *gorm.DB) *storeProcedureService {
	return &storeProcedureService{
		Db: db,
	}
}

func (s *storeProcedureService) getListFunction() error {
	var listFunction *[]seeder.StoreProcedure
	queryListSp := `select 
		pp.proname,
		pl.lanname,
		pn.nspname,
		pg_get_functiondef(pp.oid) as queryFunction
	from pg_proc pp 
		inner join pg_namespace pn on (pp.pronamespace = pn.oid) 
		inner join pg_language pl on (pp.prolang = pl.oid) 
	where pl.lanname NOT IN ('c','internal') 
		and pn.nspname NOT LIKE 'pg_%' 
		and pn.nspname <> 'information_schema';`
	res := s.Db.Raw(queryListSp).Scan(&listFunction)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("list function empty")
	} else {
		s.listFunction = listFunction
	}
	return nil
}

func (s *storeProcedureService) SyncSp() error {
	return nil
}

func (s *storeProcedureService) RunImport() error {
	return nil
}
