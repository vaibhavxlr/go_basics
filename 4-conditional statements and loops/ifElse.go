package main

import "fmt"

func main() {
	/*
	 basic "if" statement
	*/
	num := 10
	if num%2 == 0 {
		fmt.Println("The number is even")
		return
	}

	fmt.Println("The number is odd")

	/*
	 basic "if-else" statement
	*/
	num = 11

	if num%2 == 0 {
		fmt.Println("The number is even")
	} else { //else should start right where block of if
		//ends, go substitutes semicolon after
		//every line and block and not doing this way
		//will lead to error
		fmt.Println("The number is odd")
	}

	/*
		 	IDIOMATIC GO
			  In Go's philosophy, it is better to avoid
			  unnecessary branches and indentation of code.
			  It is also considered better to return as
			  early as possible.
	*/

}
