package main

import (
	"fmt"
	"goprototype/internal/shape"
)

func main() {
	rect := shape.NewRectangle(5, 5)
	newRect := rect.Clone()

	fmt.Println(&rect)
	fmt.Println(&newRect)
}
