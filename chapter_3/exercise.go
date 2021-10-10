package main

import "fmt"

func main() {
	// doubleNumber()
	// convertTemp()
	convertLength()
}

func convertLength() {
	fmt.Print("Enter a length in feet: ")
	var input float64
	fmt.Scanf("%f", &input)
	converted := input * 0.3048
	fmt.Println("Converted to meters:", converted)
}

func convertTemp() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)
	converted := convertFtoC(input)
	fmt.Println("Converted to celsius:", converted)
}

func convertFtoC(temp float64) float64 {
	res := (temp - 32) * 5 / 9
	return res
}

func doubleNumber() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("input:", input)
	output := input * 2
	fmt.Println("output", output)
}
