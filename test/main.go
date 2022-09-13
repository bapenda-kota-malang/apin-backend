package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/go-chi/chi/v5"
)

var db *gorm.DB

func main() {
	db, _ = gorm.Open(
		mysql.Open("root:12341234@tcp(127.0.0.1:3306)/sapiperah?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
		},
	)

	r := chi.NewRouter()
	r.Post("/create", create)
	r.Post("/createX", createX)
}
