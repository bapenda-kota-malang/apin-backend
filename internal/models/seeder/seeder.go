package seeder

type StoreProcedure struct {
	Proname       string `gorm:"column:proname"`
	Lanname       string `gorm:"column:lanname"`
	Nspname       string `gorm:"column:nspname"`
	Queryfunction string `gorm:"column:queryfunction"`
}

type Config interface {
	LoadConfig() error
	GenerateCommand(notrx *bool) error
}

type SeedSql struct {
	basePath string
	cmd      string
	platform string
}

type SeedSqlService interface {
	RunImportSql(cmd string) error
	GetListFile() ([]string, error)
	WriteSeedSql(files []string) error
}
