package main

import (
	"fmt"

	"github.com/zanderhavgaard/introducing-go-exercises/chapter_8/creating_packages/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
