package main

type tablex struct {
	Id   int    `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Code string `json:"code" gorm:"type:varchar"`
	Name string `json:"name" gorm:"type:varchar"`
}

type tablexInput struct {
	Code string `json:"kode"`
	Name string `json:"name"`
}
