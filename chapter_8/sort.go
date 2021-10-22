package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// create types that we can cast the Person type to when sorting
type ByName []Person
type ByAge []Person

// implement the three metods of the sort.Interface
// Len, Less, Swap

// get the length of a slice
func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByAge) Len() int {
	return len(ps)
}

// check if element i is less than element j
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

// swap positions in slice
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"foobar", 666},
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
