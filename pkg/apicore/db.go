package apicore

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
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
	AutoMigrate(&user.User{})
	AutoMigrate(&configurationmodel.Kecamatan{})
	AutoMigrate(&configurationmodel.Kelurahan{})
	AutoMigrate(&configurationmodel.Rekening{})
	AutoMigrate(&configurationmodel.Skpd{})
	AutoMigrate(&registration.Registration{})
	AutoMigrate(&registration.ObjekPajak{})
	AutoMigrate(&registration.DetailOpAirTanah{})
	AutoMigrate(&registration.DetailOpHiburan{})
	AutoMigrate(&registration.DetailOpHotel{})
	AutoMigrate(&registration.DetailOpParkir{})
	AutoMigrate(&registration.DetailOpPpj{})
	AutoMigrate(&registration.DetailOpReklame{})
	AutoMigrate(&registration.DetailOpResto{})
	AutoMigrate(&registration.PemilikWp{})
	AutoMigrate(&registration.Narahubung{})
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

	for _, item := range autoMigrateList {
		DB.AutoMigrate(&item)
	}
}

func AutoMigrate(model interface{}) {
	autoMigrateList = append(autoMigrateList, model)
}
