package dbmigrator

type Migrator struct {
	Src Db `json:"src"`
	Dst Db `json:"dst"`
}

type Db struct {
	Dialect string `json:"dialect"`
	Table   Table  `json:"table"`
}

type Table struct {
	Name   string   `json:"name"`
	Column []string `json:"column"`
}

// type DbMigrator struct {
// 	Data []Migrator `json:"data"`
// }

type DbMigratorService interface {
	ParseJson(b []byte) error
	RunCopyData() error
}
