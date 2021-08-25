package main

import "fmt"

func main() {
	var width, height int = 30, 200 //declaring multiple variables
	fmt.Println("width is", width, "height is", height)

	var w, h = 20.0, 30.9 //dropped data type
	fmt.Println("width is", w, "height is", h)

	w, h = 1, 2
	fmt.Println("width is", w, "height is", h)

	var (
		name        = "vaibhav"
		age         = "21"
		carryHeight int
	)
	carryHeight = 172
	fmt.Println("name is", name, "age is", age, "height is", carryHeight)

}
