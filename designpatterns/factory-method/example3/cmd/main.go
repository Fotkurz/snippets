package main

import "golangsimplefactory/internal"

func main() {
	store := internal.NewPizzaStore()

	pepperoni := store.Order("pepperoni")
	cheese := store.Order("cheese")
	random := store.Order("i don't know")

	println(pepperoni.Serve())
	println(cheese.Serve())
	println(random.Serve())
}
