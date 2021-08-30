package main

import "fmt"

func main() {

	var age int                    //variable declaration
	fmt.Println("My age is ", age) //uninitialised variables are assigned with default value zero

	age = 10
	fmt.Println("My age is ", age)

	age++
	fmt.Println("My age is ", age)

	var superAge int = 20 //varibale with initial value
	fmt.Println("My age is ", superAge)
}
