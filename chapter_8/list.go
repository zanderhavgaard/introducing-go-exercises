package main

import (
	"container/list"
	"fmt"
)

func main() {
	// create linked list
	var myList list.List

	// add data to list
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	// you are allowed to mix types in a list
	myList.PushBack("foobar")
	myList.PushBack(5)

	// iterate over list items and print them
	for e := myList.Front(); e != nil; e = e.Next() {
		// will just print the value
		fmt.Println(e.Value)
		// access the value, but typecast as integer
		// does not work with mixed types for obvious reasons...
		// fmt.Println(e.Value.(int))
	}
}
