package seeder

import "flag"

type Option struct {
	Notrx     *bool
	UpdateSql *bool
}

func NewOption() *Option {
	cmd := defineFlag()
	flag.Parse()
	return cmd
}

// get flag and assign to struct
func defineFlag() *Option {
	return &Option{
		Notrx:     flag.Bool("notrx", false, "for using transaction when execute seeder psql command"),
		UpdateSql: flag.Bool("updatesql", false, "for auto update list seed.sql with sql file inside sqls folder"),
	}
}
