package main

import "fmt"

//example function 1
func func1(a int, b int) int {
	return a * b
}

//example function 2
func func2(a, b int) int {
	return a * b
}

//example function 3
func func3(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

//example function 4
func func4(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}

//example function 5
func func5(a, b int) {
	fmt.Println("\nanswer is:", a+b)
}

func main() {
	fmt.Printf("func1 gives %v", func1(2, 2))
	fmt.Printf("\nfunc2 gives %v", func2(2, 2))

	val1, val2 := func3(2, 2)
	fmt.Printf("\nfunc3 gives %v and %v", val1, val2)

	val1, val2 = func4(2, 2)
	fmt.Printf("\nfunc4 gives %v and %v", val1, val2)

	func5(2, 2)

	val1, _ = func3(3, 3) //second value not required so used BLANK IDENTIFIER
	fmt.Println("val1 is:", val1)

}
