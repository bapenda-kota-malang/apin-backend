package seeder

import "flag"

type option struct {
	Notrx          *bool
	UpdateSql      *bool
	StoreProcedure *string
}

func NewOption() *option {
	return &option{}
}

func (o *option) ParseOption() {
	o.Notrx = flag.Bool("notrx", false, "for using transaction when execute seeder psql command")
	o.UpdateSql = flag.Bool("updatesql", false, "for auto update list seed.sql with sql file inside sqls folder")
	o.StoreProcedure = flag.String("sp", "skip", "for import store procedure, sync for read first then run or run for run only without sync from db")
	flag.Parse()
	return
}
