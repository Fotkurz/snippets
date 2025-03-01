package main

import (
	"abstractfactoryexample1/internal/furniture"
	"log"
)

func main() {
	factory, err := furniture.NewFurnitureFactory(furniture.ModernFurnitureType)
	if err != nil {
		log.Fatal(err)
	}

	printFurnitures(factory)

	factory, err = furniture.NewFurnitureFactory(furniture.VictorianFurnitureType)
	if err != nil {
		log.Fatal(err)
	}

	printFurnitures(factory)

	_, err = furniture.NewFurnitureFactory("random")
	if err != nil {
		log.Fatal(err)
	}

}

func printFurnitures(factory furniture.FurnitureFactory) {
	chair := factory.CreateChair()
	table := factory.CreateTable()
	sofa := factory.CreateSofa()

	println(chair.TypeName() + " chair")
	println(table.TypeName() + " table")
	println(sofa.TypeName() + " sofa")

}
