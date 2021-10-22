package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// setup the cli argument
	max := flag.Int("max", 6, "the max value")
	// parse cli arguments
	flag.Parse()
	//Generate a random number between 0 and max
	num := rand.Intn(*max)
	fmt.Println(num)
}

/*
 * Run the program with: $ go run cli.go -max=100
 * To set the argument
 */
