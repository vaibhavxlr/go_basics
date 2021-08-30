package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 50
	fmt.Println(a)

	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name, age, country)

	const x = 66 //allowed
	x = 86       // reassignment not allowed

	var y = math.Sqrt(5)   //allowed
	const z = math.Sqrt(5) //not allowed, since the assignment
	//happens at runtime which is not possible
	//because it is a constant

	const hello1 = "Hello world"             //untyped string constant
	const typedHello1 string = "Hello world" //typed string constant

	//One way to think about untyped constants is that they
	//live in a kind of ideal space of values, a space less
	//restrictive than Go’s full type system. But to do
	//anything with them, we need to assign them to variables,
	//and when that happens the variable (not the constant
	//itself) needs a type, and the constant can tell the
	//variable what type it should have. In this example,
	//str becomes a value of type string because the untyped
	//string constant gives the declaration its default type, string.
	const sam = "Sam"
	var varSam = sam
	fmt.Printf("type %T value %v", varSam, varSam) //output:- type string value Sam

	//**following has nothing to do with typed and untyped as
	//**they are for constants, although similar properties
	//**will help understand typed and untyped constants better
	var defaultName = "Sam"         //allowed
	type myString string            //aliasing
	var customName myString = "Sam" //allowed
	customName = defaultName        //not allowed, because although both are fundamentally strings
	//but different types of strings(one is string and another is myString),
	//Go compiler considers them of different types

	//boolean constants
	const trueConst = true
	var defaultbool = trueConst //allowed
	type mybool bool
	var custombool mybool = trueConst //allowed
	defaultbool = custombool          //not allowed

	//numeric constants
	var i = 5
	var j = 5.5
	var k = 5 + 6i
	fmt.Printf("i is of type %T, j is of type %T, k is of type %T", i, j, k)
	//output:- i is of type int, j is of type float64, c is of type complex128
	//therefore, variables assume the type by the values supplied to them
	//untyped constants act in similar way, they are like values with implicit types

	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Printf("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	//output:- 	  intVar 5
	//    int32Var 5
	//    float64Var 5
	//    complex64Var (5 + 0i)
	//therefore untyped constants assume their default type depending upon context and hence
	//act as values supplied above in the line 62 to line 64

	//example to reiterate what is written in line 79 and 80
	const X1 = 5.9
	const Y1 = 8
	var a = X1 / Y1
	var b = 5.9 / 8

	if a == b {
		fmt.Println("True")
	}
	//output:- True
}
