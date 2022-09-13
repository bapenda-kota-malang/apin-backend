package main

import "fmt"

func createData(data any, dataInput any) {

}

func createDataX(data tablex, dataInput tablexInput) {
	db.Create(&data)
	fmt.Print()
}
