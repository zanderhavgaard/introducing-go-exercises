package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readFile1() {
	file, err := os.Open("io.go")
	if err != nil {
		// if error is present, then handle it
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// fmt.Println("Size of file", stat)

	// read the file to bytes
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func readFile2() {
	// much less verbose and better way to read contents of file
	filename := "io.go"
	rawContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contents := string(rawContents)
	fmt.Println(contents)
}

func writeFile() {
	// this will overwrite any file with the same name!
	file, err := os.Create("created-by-golang.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// write to the file
	file.WriteString("foobar\n")
	file.WriteString("barfoo")
}

func dirContents() {
	// open current directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	// how many files to get info on, -1 for all
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	// iterate over file infos and print the name of each file
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func recDirContents1() {
	// recursively walk the directory tree and print all file/dir names
	// uses a lambda function to print the path to each dir/file encountered
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		// same as above
		// fmt.Println(info.Name())

		// if no error then return nil
		return nil
	})
}

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

func recDirContents2() {
	filepath.Walk(".", walkFunc)
}

func recDirContents3() {
	f := func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	}
	filepath.Walk(".", f)
}

func main() {
	// readFile1()
	// readFile2()
	// writeFile()
	// dirContents()
	recDirContents1()
	recDirContents2()
	recDirContents3()
}
