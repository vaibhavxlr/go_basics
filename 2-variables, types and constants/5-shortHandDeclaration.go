package main

import "fmt"

func main() {
	count := 10 //short hand operator -> :=
	fmt.Println("count", count)

	name, age := "vaibhav", 21 //short hand with multiple variable declaration
	fmt.Println("name is", name, "age is", age)
}
