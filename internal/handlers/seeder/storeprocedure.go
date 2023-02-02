package seeder

import "gorm.io/gorm"

type seedSp struct {
	Db *gorm.DB
}

func NewSeedSp(db *gorm.DB) *seedSp {
	return &seedSp{
		Db: db,
	}
}

func (s *seedSp) SyncSp() error {
	return nil
}

func (s *seedSp) RunImport() error {
	return nil
}
