package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		bool
	*/
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)

	/*
		signed and unsigned integers
	*/

	var e int = 88
	f := 99
	var g uint32 = 999

	fmt.Printf("e is of type %T, e is of size %d", e, unsafe.Sizeof(e))
	fmt.Printf("\nf is of type %T, f is of size %d", f, unsafe.Sizeof(f))
	fmt.Printf("\ng is of type %T, g is of size %d", g, unsafe.Sizeof(g))

	/*
		floating point numbers
	*/

	var h float32 = 88.88
	i := 99.098
	var j float64 = 999.88

	fmt.Printf("\nh is of type %T, h is of size %d", h, unsafe.Sizeof(h))
	fmt.Printf("\ni is of type %T, i is of size %d", i, unsafe.Sizeof(i))
	fmt.Printf("\nj is of type %T, j is of size %d", j, unsafe.Sizeof(j))

	/*
		string type
	*/

	name := "Vaibhav"
	age := "21"

	fmt.Println("\n" + name + ", " + age)

	/*
		type conversion

		Trying to add 2 numbers of
		different types is not allowed.
		Explicit type conversion is required
		to assign a variable of one type to
		another.
	*/

	x := 75
	y := 64.4

	sum := x + int(y)
	fmt.Println(sum)

	z1 := 75

	var z2 float32 = float32(z1) //won't work without explicit conversion
	fmt.Printf("\nz2 is:%f, z1 is:%d, z2 is of type:%T, z1 is of type:%T", z2, z1, z1, z2)
}
