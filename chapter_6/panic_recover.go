package main

import "fmt"

func doesNotWork() {
	// this panic will never be recovered
	panic("panic")
	str := recover()
	fmt.Println(str)
}

func main() {
	defer func() {
		// recover will recover from a runtime error
		// and return the message from the panic
		str := recover()
		fmt.Println(str)
	}()
	// panic triggers a runtime error
	panic("PANIC")
}
