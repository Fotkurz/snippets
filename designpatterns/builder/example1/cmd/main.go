package main

import (
	"builderexample1/internal/car"
	"fmt"
)

func main() {

	builder := car.NewCarAutomaticBuilder()
	director := car.Director{}

	car := director.MakeAutomaticCar(builder)

	fmt.Printf("%v", car)
}
