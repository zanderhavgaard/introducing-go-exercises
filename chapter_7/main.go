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

// method, that only works on Circle structs
// the (c *Circle) is called the 'reciever' in a method signature
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// define type for rectangles
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// define method for rectangle struct to calculate area
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// define person type
type Person struct {
	Name string
}

// method on Person struct
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// struct with embedded type, an Android is a Person
// sometimes also called an 'anonymous filed'
type Android struct {
	Person
	Model string
}

// create interface for shapes, structs must then implement the interface,
// in order to be able to use the interface methods
type Shape interface {
	area() float64
}

// create interface method for the shape interface
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// alternate implicit syntax
func totalArea2(shapes ...Shape) (area float64) {
	for _, s := range shapes {
		area += s.area()
	}
	return
}

// a struct that uses an interface
type MultiShape struct {
	Shapes []Shape
}

// implement the area() method for the MultiShape struct,
// which makes the MultiShape a Shape, as it implements the interface
func (m *MultiShape) area() (area float64) {
	for _, s := range m.Shapes {
		area += s.area()
	}
	return
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

	area := c.area()
	fmt.Println(area)

	rec := Rectangle{0, 0, 10, 10}
	area2 := rec.area()
	fmt.Println(area2)

	person := Person{"foo"}
	person.Talk()

	robot := new(Android)
	fmt.Println(robot)
	// we can call methods of the embedded Person type
	robot.Person.Talk()
	// because Person is embedded we can call them directly
	robot.Talk()

	// declare struct variable with embedded type
	robot2 := Android{Person: Person{Name: "bar"}, Model: "123"}
	fmt.Println(robot2)
	robot2.Person.Talk()
	robot2.Talk()

	// use interface method
	ta := totalArea(&c, &rec)
	fmt.Println(ta)

	// declare MultiShape struct using pointer types
	ms := MultiShape{
		Shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	area3 := ms.area()
	fmt.Println(area3)
	// for how to implement this without pointer types, see the book errata: https://www.oreilly.com/catalog/errata.csp?isbn=0636920046516
}
