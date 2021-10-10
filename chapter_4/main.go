package main

import "fmt"

func main() {
	// golang only has 'for' loops, but it can be used in a number of different ways

	// verbose syntax
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// more concise syntax
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	oddOrEven()
	printStringOfNumber()
}

func printStringOfNumber() {
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			{
				fmt.Println("zero")
				fmt.Println("zero")
				fmt.Println("zero")
			}
		case 1:
			fmt.Println("one")
			fmt.Println("one")
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		}
	}
}

func oddOrEven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else if i%3 == 0 {
			fmt.Println(i, "is divisible by 3")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
