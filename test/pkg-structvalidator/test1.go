package main

import (
	"fmt"

	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
)

type CarModel string
type ChairModel string

const (
	Truct CarModel = "truct"
	Sedan CarModel = "sedan"
	Suv   CarModel = "suv"
	Bus   CarModel = "bus"
)

const (
	FixedChair   ChairModel = "fixedChair"
	DynamicChair ChairModel = "dynamicChairchair"
)

type Dimension struct {
	Width  int `validate:"required;min=1"`
	Height int `validate:"required;min=1"`
	Length int `validate:"required;min=1"`
}

type chair struct {
	Model ChairModel `validate:"required"`
	Dimension
}

type car struct {
	Model  CarModel `validate:"required"`
	Base64 string   `validate:"required;base64"`
	Dimension
	Chair []ChairModel
}

func main() {
	myCar := car{
		Model:  Bus,
		Base64: "1234=a",
		Dimension: Dimension{
			Width: 2,
		},
	}
	myCar.Height = 100
	x := sv.Validate(myCar)
	fmt.Println(x)
}
