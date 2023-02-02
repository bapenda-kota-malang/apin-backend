package seeder

type StoreProcedure struct {
	Proname       string
	Lanname       string
	Nspname       string
	QueryFunction string
}

type Config interface {
	LoadConfig() error
	GenerateCommand(notrx *bool) error
}
