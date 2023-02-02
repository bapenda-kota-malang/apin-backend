package seeder

type Config interface {
	LoadConfig() error
	GenerateCommand(notrx *bool) error
}
