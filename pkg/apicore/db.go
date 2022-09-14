package apicore

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type dbConf struct {
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
	Dialect      string
}

var autoMigrateList []interface{}

func init() {
	// autoMigrateList = make([]interface{})
	listModelPendaftaran := []interface{}{
		&user.User{},
		&configurationmodel.Provinsi{},
		&configurationmodel.Daerah{},
		&configurationmodel.Kecamatan{},
		&configurationmodel.Kelurahan{},
		&configurationmodel.Rekening{},
		&configurationmodel.Skpd{},
		&registration.Registration{},
		&registration.ObjekPajak{},
		&registration.DetailOpAirTanah{},
		&registration.DetailOpHiburan{},
		&registration.DetailOpHotel{},
		&registration.DetailOpParkir{},
		&registration.DetailOpPpj{},
		&registration.DetailOpReklame{},
		&registration.DetailOpResto{},
		&registration.PemilikWp{},
		&registration.Narahubung{},
	}
	AutoMigrate(listModelPendaftaran...)

	listModelPendataan := []interface{}{
		&potensiopwp.PotensiOp{},
		&potensiopwp.PotensiPemilikWp{},
		&potensiopwp.PotensiNarahubung{},
		&potensiopwp.DetailPotensiOp{},
		&potensiopwp.DetailPotensiAirTanah{},
		&potensiopwp.DetailPotensiHiburan{},
		&potensiopwp.DetailPotensiHotel{},
		&potensiopwp.DetailPotensiPPJ{},
		&potensiopwp.DetailPotensiParkir{},
		&potensiopwp.DetailPotensiReklame{},
		&potensiopwp.DetailPotensiResto{},
	}
	AutoMigrate(listModelPendataan...)
}

func (a *app) initDb() {
	if a.DbConf.Dsn == "" {
		Logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "unset"),
			zap.String("message", "no DSN provided"))
		return
	}
	if a.DbConf.Dialect != "mysql" && a.DbConf.Dialect != "postgres" {
		Logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "unset"),
			zap.String("message", "no proper dialect provided"))
		return
	}

	var gormD gorm.Dialector
	if a.DbConf.Dialect == "mysql" {
		gormD = mysql.Open(a.DbConf.Dsn)
	} else if a.DbConf.Dialect == "postgres" {
		gormD = postgres.Open(a.DbConf.Dsn)
	}

	db, err := gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		Logger.Panic("instantiation", zap.String("type", "db"), zap.String("source", "gorm"), zap.String("status", "panic"))
		panic("Failed to connect to database!")
	} else {
		DB = db
		Logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "done"),
			zap.String("name", "db"))
	}

	DB.AutoMigrate(autoMigrateList...)
}

func AutoMigrate(model ...interface{}) {
	autoMigrateList = append(autoMigrateList, model...)
}
