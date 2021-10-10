package main

import "fmt"

func main() {
	// smallBig()
	// divisibleByThree()
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func divisibleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func smallBig() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}
