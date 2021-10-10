package main

import "fmt"

func main() {
	// arrayExample1()
	// arrayExample2()
	// arrayExample3()
	// sliceExample1()
	// sliceExample2()
	// sliceExample3()
	// mapExample1()
	// mapExample2()
	// mapExample3()
	mapExample4()
}

func arrayExample1() {
	// declare and instantiate array of ints
	// arrays are of a static size
	var myArr [5]int
	// arrays are  indexed from zero
	myArr[0] = 1
	myArr[4] = 100
	fmt.Println(myArr)
	fmt.Println(myArr[4])
}

func arrayExample2() {
	const arrSize int = 5
	var myFloatArr [arrSize]float64
	myFloatArr[0] = 98
	myFloatArr[1] = 93
	myFloatArr[2] = 77
	myFloatArr[3] = 82
	myFloatArr[4] = 83

	var total float64 = 0
	for i := 0; i < len(myFloatArr); i++ {
		total += myFloatArr[i]
	}

	fmt.Println(total / float64(len(myFloatArr)))
}

func arrayExample3() {
	var myFloatArr [5]float64
	myFloatArr[0] = 98
	myFloatArr[1] = 93
	myFloatArr[2] = 77
	myFloatArr[3] = 82
	myFloatArr[4] = 83

	var total float64 = 0
	for _, value := range myFloatArr {
		total += value
	}

	fmt.Println(total / float64(len(myFloatArr)))
}

func arrayExample4() {
	myArr1 := [5]float64{98, 93, 77, 82, 83}
	myArr2 := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println(myArr1)
	fmt.Println(myArr2)
}

func sliceExample1() {
	// declare and instantiate an empty slice, but without an associated array
	var mySlice []float64
	// mySlice[0] = 1 -- this does not work, as there is no array to enter data into

	fmt.Println(mySlice)

	// use built-in function to create slice with assiciated array
	mySlice2 := make([]float64, 5)
	// works
	mySlice2[0] = 1
	// is out of bounds
	// mySlice2[5] = 1

	fmt.Println(mySlice2)

	// create a slice and associate array
	mySlice3 := make([]float64, 5, 10)
	mySlice3[1] = 1

	fmt.Println(mySlice3)

	// create array
	arr := [5]float64{1, 2, 3, 4, 5}
	// create slice from array indexes [<low index>:<high index>]
	// the high index is not included, so high-index of 5 points to 4, which is the fifth element
	mySlice4 := arr[0:5]
	fmt.Println(mySlice4)
	// you may omit low or high index
	// this will use 0 or the len(arr) respectively
	fmt.Println(arr[0:])
	fmt.Println(arr[2:])
	fmt.Println(arr[:5])
	fmt.Println(arr[:3])
}

func sliceExample2() {
	// this is a slice b/c the array does not have a size argument
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)
	// appending to a slice beyond it's capacity will create a new slice
	// (with a new underlying array), which is returned
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(slice2)
}

func sliceExample3() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	fmt.Println(slice1, slice2)
	numItemsCopied := copy(slice2, slice1)
	fmt.Println(slice1, slice2, numItemsCopied)
}

func mapExample1() {
	// declare a map with string keys and int values
	var aMap map[string]int
	fmt.Println(aMap)
	// declare and instantiate a map
	myMap := make(map[string]int)
	// alternative syntax
	// var myMap = make(map[string]int)
	// var myMap map[string]int = make(map[string]int)
	myMap["myKey"] = 10
	fmt.Println(myMap)
	fmt.Println(myMap["myKey"])
	fmt.Println(myMap["not a used key"])

	myIntMap := make(map[int]int)
	myIntMap[1] = 10
	fmt.Println(myIntMap[1])
	fmt.Println(myIntMap[10])
	fmt.Println(myIntMap, len(myIntMap))
	// remove a key from a map
	delete(myIntMap, 1)
	fmt.Println(myIntMap, len(myIntMap))
}

func mapExample2() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	// print name if key exists
	name, ok := elements["Un"]
	if ok {
		fmt.Println(name)
	}

	// more elegant syntax
	if name, ok := elements["Un"]; ok {
		fmt.Println(ok, name)
	}
}

func mapExample3() {
	// declare, instantiate and fill a map
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func mapExample4() {
	// nested maps
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// iterate over map
	for key, val := range elements {
		fmt.Println(key)
		for k2, v2 := range val {
			fmt.Println(" ", k2, v2)
		}
	}
}
