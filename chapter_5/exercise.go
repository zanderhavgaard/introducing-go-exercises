package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest int
	for _, i := range x {
		fmt.Println(i)
		if smallest == 0 {
			smallest = i
		} else if i < smallest {
			smallest = i
		}
	}
	fmt.Println("Smallest int", smallest)
}
