package main

import "errors"

func main() {
	err := errors.New("error message")

	// throw the error with panic
	panic(err)
}
