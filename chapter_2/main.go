package main

import "fmt"

func main() {
	// numbers: int + float
	fmt.Println("1+1=", 1+1)
	fmt.Println("1.0+1.0=", 1.0+1.0)

	// strings
	fmt.Println(len("Hello world!"))
	fmt.Println(string("Hello world"[0]))
	fmt.Println("hello," + " " + "world")

	// booleans
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// Exercise 3
	fmt.Println(32,132 * 42.452)

	// Exercise 5
	fmt.Println((true && false) || (false && true) || !(false && false))
}
