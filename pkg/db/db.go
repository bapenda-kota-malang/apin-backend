package db

import (
	"fmt"

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

func NewDbConf(dsn string, dialect string) *dbConf {
	return &dbConf{Dsn: dsn, Dialect: dialect}
}

func (c *dbConf) GenerateDsn(host string, dbUser string, password string, dbName string, port string) error {
	if host == "" {
		return fmt.Errorf("host empty")
	}
	if dbUser == "" {
		return fmt.Errorf("dbUser empty")
	}
	if port == "" {
		return fmt.Errorf("port empty")
	}
	if password == "" {
		return fmt.Errorf("password empty")
	}
	if dbName == "" {
		return fmt.Errorf("dbName empty")
	}
	c.Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, dbUser, password, dbName, port)
	return nil
}

func (c *dbConf) InitDb() (db *gorm.DB, err error) {
	if c.Dsn == "" {
		err = fmt.Errorf("dsn not provided")
		return
	}
	if c.Dialect != "mysql" && c.Dialect != "postgres" {
		err = fmt.Errorf("no proper dialect provided")
		return
	}

	var gormD gorm.Dialector
	if c.Dialect == "mysql" {
		gormD = mysql.Open(c.Dsn)
	} else if c.Dialect == "postgres" {
		gormD = postgres.Open(c.Dsn)
	}

	db, err = gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		return
	}
	return
}
