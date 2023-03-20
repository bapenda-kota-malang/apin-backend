package seeder

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/seeder"
	"gorm.io/gorm"
)

type seedSp struct {
	Db *gorm.DB
}

func NewSeedSp(db *gorm.DB) *seedSp {
	return &seedSp{
		Db: db,
	}
}

// get list function from db
func (s *seedSp) GetListFunction() ([]seeder.StoreProcedure, error) {
	var listFunction []seeder.StoreProcedure
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
	resRows, err := s.Db.Raw(queryListSp).Rows()
	if err != nil {
		return nil, err
	}
	defer resRows.Close()

	for resRows.Next() {
		var spRow seeder.StoreProcedure
		s.Db.ScanRows(resRows, &spRow)
		listFunction = append(listFunction, spRow)
	}

	if len(listFunction) == 0 {
		return nil, fmt.Errorf("list function empty")
	}

	return listFunction, nil
}
