package main

import (
	"fmt"
)

// with return type
func average(scores []float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

// with named return value
func average2(scores []float64) (total float64) {
	total = 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

// return muliple values
func multiReturn() (int, int) {
	return 1, 2
}

// 'variadic' function argument
// when invoking function any number of int arguments can be given
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// create a function that returns a 'closure' function
// the returned function will have it's own local 'i' variable
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursive function that computes factorial
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}

func main() {
	scores := []float64{98, 93, 77, 82, 83}
	result := average(scores)
	fmt.Println(result)

	// assign multiple variables from return values from function
	a, b := multiReturn()
	fmt.Println(a, b)

	added := add(1, 2, 3, 4)
	added = add(1, 2, 3, 4, 5, 6)
	fmt.Println(added)

	// a slice can be given as a variadic argument, by appending an 'ellipsis'
	nums := []int{1, 2, 3, 4}
	numsAdded := add(nums...)
	fmt.Println(numsAdded)

	//declare function in scope of another function
	add2 := func(x, y int) int {
		return x + y
	}
	added2 := add2(2, 3)
	fmt.Println(added2)

	//local functions have access to local scope
	x := 0
	inc := func() {
		x++
	}
	inc()
	fmt.Println(x)
	inc()
	inc()
	fmt.Println(x)

	y := 0
	inc2 := func() int {
		y++
		return y
	}
	fmt.Println(inc2())
	fmt.Println(inc2())
	// we call these types of functions 'closures'

	// use closure function to create local function, with local scope
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	factorial2 := factorial(10)
	fmt.Println(factorial2)

	// defer keyword makes the function call be called after the wrapping function completes
	// thus in this example second will be executed after main completes
	defer second()
	first()
	// this is really need for things like making sure a file stream is closed after we are done with it
	// example:
	// f, _ := os.Open(filename)
	// defer f.Close()

}
