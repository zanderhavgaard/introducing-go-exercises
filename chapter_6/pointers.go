// pointers lets us refer to a specific variable
// from a different scope, for example in a function
// normally the value of a variable would be copied as an argument
// to the function, but by using a pointer, we can interact
// with the specific variable, but from the function scope
package main

import "fmt"

// does not use a pointer, copies the value of x
func zero(x int) {
	x = 0
}

// take pointer to x as arg
// the asterisk dereferences a pointer
func zeroPointer(x *int) {
	// modify the variable the pointer points to
	*x = 0
}

func main() {
	x := 5

	// this does nothing as the x in this scope is coped to zero()
	zero(x)
	fmt.Println(x)

	// here we send a pointer to x as the argument
	// the local variable x is modified in the zeroPointer() function
	// b/c it uses the pointer to the local x variable
	// the ampersand gives us the memory location of the x variable
	// &x -> *x the ampersand gives us a pointer to the x variable in memory
	// which is then dereferenced with the asterisk
	zeroPointer(&x)
	fmt.Println(x)
}
