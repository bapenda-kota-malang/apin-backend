package apicore

import (
	adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	rnpwpd "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
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
		&wajibpajak.WajibPajak{},
		&adm.Provinsi{},
		&adm.Daerah{},
		&adm.Kecamatan{},
		&adm.Kelurahan{},
		&rm.Rekening{},
		&skpd.Skpd{},
		&npwpd.Npwpd{},
		&npwpd.ObjekPajak{},
		&npwpd.DetailOpAirTanah{},
		&npwpd.DetailOpHiburan{},
		&npwpd.DetailOpHotel{},
		&npwpd.DetailOpParkir{},
		&npwpd.DetailOpPpj{},
		&npwpd.DetailOpReklame{},
		&npwpd.DetailOpResto{},
		&npwpd.PemilikWp{},
		&npwpd.Narahubung{},
		&rnpwpd.RegistrasiNpwpd{},
		&rnpwpd.RegObjekPajak{},
		&rnpwpd.DetailRegOpAirTanah{},
		&rnpwpd.DetailRegOpHiburan{},
		&rnpwpd.DetailRegOpHotel{},
		&rnpwpd.DetailRegOpParkir{},
		&rnpwpd.DetailRegOpPpj{},
		&rnpwpd.DetailRegOpReklame{},
		&rnpwpd.DetailRegOpResto{},
		&rnpwpd.RegPemilikWp{},
		&rnpwpd.RegNarahubung{},
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
