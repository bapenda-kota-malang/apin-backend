package main

import (
	"fmt"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request) {
	var data tablex
	var input tablexInput
	getPayload(&input, r.Body)
	fmt.Println(data)
	createData(data, input)
}

func createX(w http.ResponseWriter, r *http.Request) {
	var data tablex
	var input tablexInput
	getPayload(&input, r.Body)
	fmt.Println(data)
	createDataX(data, input)
}
