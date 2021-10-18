package main

import (
	"fmt"
	"strings"
)

func main() {
	// check if substring
	fmt.Println(strings.Contains("test", "es"))

	// count how many times substring in string
	fmt.Println(strings.Count("test", "t"))

	// check string starts with substring
	fmt.Println(strings.HasPrefix("test", "te"))

	// check if string ends with substring
	fmt.Println(strings.HasSuffix("test", "st"))

	// find position of a substring in string
	fmt.Println(strings.Index("test", "e"))

	// take a list of strings and join them to a sinlge string
	// second argument is but in between each string to be joined
	joined := strings.Join([]string{"a", "b"}, "-")
	fmt.Println(joined)

	// repeat a string n number of times
	fmt.Println(strings.Repeat("foo", 3))

	// replace substring with another substring in string and number of occurence of substring to replace
	fmt.Println(strings.Replace("foobarfoobar", "foo", "bar", 1))
	// to replace all occurances use -1
	fmt.Println(strings.Replace("foobarfoobar", "foo", "bar", -1))

	// split string by each occurance of substring
	var splitSlice []string
	splitSlice = strings.Split("a,b,c,d,e", ",")
	fmt.Println(splitSlice)

	// convert to lowercase
	fmt.Println(strings.ToLower("FOOBAR"))

	// convert to uppercase
	fmt.Println(strings.ToUpper("foobar"))

	// convert string to slice of bytes and back to string
	arr := []byte("test")
	str := string([]byte{'a', 'b'})
	fmt.Println(arr)
	fmt.Println(string(arr))
	fmt.Println(str)
}
