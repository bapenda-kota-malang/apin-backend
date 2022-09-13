package apicore

import (
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
