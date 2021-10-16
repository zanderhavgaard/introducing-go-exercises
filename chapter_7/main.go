package main

import (
	"fmt"
	"math"
)

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, x2, y1, y2 float64) float64 {
	l := distance(x1, x2, y1, y2)
	w := distance(x1, x2, y1, y2)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

// define circle struct
type Circle struct {
	x, y, r float64
}

// function that takes a struct as argument
func circleArea2(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// function that takes a pointer to a struct as argument
func circleArea3(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// declare assign a new struct variable
	// all fields have a zero-values
	var c1 Circle
	fmt.Println(c1)

	// use the new() function to declare and assign a circle struct
	// as well as a pointer to the struct
	// this is an uncommon usecase though
	c2 := new(Circle)
	fmt.Println(c2)

	// create new struct and assign each field by name
	c3 := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c3)

	// without field names
	c4 := Circle{0, 0, 5}
	fmt.Println(c4)

	// create a Circle struct and return the pointer
	c5 := &Circle{0, 0, 5}
	fmt.Println(c5)

	fmt.Println("---")

	c := Circle{0, 0, 5}

	// we use .<field name> to access the fields
	fmt.Println(c.x, c.y, c.r)

	// call function using struct
	fmt.Println(circleArea2(c))
	// call funciton using pointer to struct
	fmt.Println(circleArea3(&c))
}
