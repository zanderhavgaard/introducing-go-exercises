package main

import "fmt"

// this variable has global scope b/c it is declared in the global scope
// each narrowing of scope is done with curly braces {}
// each narrower scope inherits the scope of its parent
var aStr = "this string has global scope"

// aStr := "why is this not allowed?"
// TODO tried to declare this variable with walrus operator, but was not allowed
// why is that?

const myConst string = "this is a immutable string"

func main() {

	fmt.Println(myConst)
	// myConst = "this is illegal"

	// declaration and assignment
	var myString string = "Hello, World"
	fmt.Println(myString)
	myString = "world, hello"
	fmt.Println(myString)

	// implicit type declaration and assignment
	var myImplicitString = "Hello, World"

	// explicit declaration and assignment
	var myOtherString string
	myOtherString = "A different string"

	fmt.Println(myImplicitString)
	fmt.Println(myOtherString)

	// assignment
	var firstAndSecond string
	firstAndSecond = "first"
	fmt.Println(firstAndSecond)
	firstAndSecond = firstAndSecond + " second"
	fmt.Println(firstAndSecond)
	firstAndSecond += " third"
	fmt.Println(firstAndSecond)

	var str1 string = "hello"
	var str2 string = "world"
	fmt.Println(str1 == str2)
	var str3 string = "hello"
	fmt.Println(str1 == str3)

	// declare and assign variable with inferred type
	str4 := "walrus operator"
	fmt.Println(str4)
	var str5 = "this is neat"
	fmt.Println(str5)

	// naming variables
	dogsName := "Merlin"
	// golang seems to favor camel case
	dogs_name := "Merlin"
	// though snake case is also possible
	fmt.Println(dogsName, dogs_name)

	printGlobalString()
	declareMultipleVars()
}

func printGlobalString() {
	fmt.Println(aStr)
}

func declareMultipleVars() {
	// decalring mulitple variables or constants in one command
	// looks like type is inferred, or can be added explicitly
	var (
		a int    = 1
		b string = "2"
		c        = 3
	)
	fmt.Println(a, b, c)

	const (
		x = 1
		y = 2
		z = 3
	)
	fmt.Println(x, y, z)
}
