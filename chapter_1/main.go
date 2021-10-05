package main

import (
	"fmt"
	"os"
)

func main() {
	hello("golang")
	os.Exit(0)
}

func hello(name string) {
	fmt.Println("hello", name, "!")
}
