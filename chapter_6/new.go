package main

import "fmt"

func one(xPointer *int) {
	*xPointer = 1
}

func main() {
	// new(<type) allocates just enough memory for a variable of that type
	// and returns a pointer to it
	xPointer := new(int)
	one(xPointer)
	fmt.Println(*xPointer)
}
