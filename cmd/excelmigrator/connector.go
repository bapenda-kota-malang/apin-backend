package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	var err error

	errENV := godotenv.Load("config.env")
	if errENV != nil {
		log.Fatalf("error loading env file")
	}

	config := struct {
		Dsn     string
		Dialect string
	}{
		Dsn:     os.Getenv("Dsn"),
		Dialect: os.Getenv("DBdialect"),
	}

	var gormD gorm.Dialector
	if config.Dialect == "mysql" {
		gormD = mysql.Open(config.Dsn)
	} else if config.Dialect == "postgres" {
		gormD = postgres.Open(config.Dsn)
	}

	DB, err = gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
}
