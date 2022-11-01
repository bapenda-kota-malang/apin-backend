package main

import "fmt"

func main() {
	InitDB()

	fmt.Println("start importing")
	// ImportRekening()
	ImportRegion()
}
