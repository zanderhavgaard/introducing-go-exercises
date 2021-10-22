package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func hashExample1() {
	// create a hasher
	hasher := crc32.NewIEEE()

	// create some data to hash
	data := []byte("test")

	// write data to the hasher
	// hashers implement the Writer interface,
	// so we can write to them like any other Writer
	hasher.Write(data)

	// caluclate the crc32 checksum hash of the data
	hashed := hasher.Sum32()
	fmt.Println(hashed)
}

func getHash(filename string) (uint32, error) {
	// open the specified file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// close opened file
	defer file.Close()

	// create hasher
	hasher := crc32.NewIEEE()

	// copy the file contents to the hasher
	// io.Copy takes two arguments: (destination, source), returns (bytesWritten, error)
	_, err = io.Copy(hasher, file)
	if err != nil {
		return 0, err
	}

	return hasher.Sum32(), nil

}

func hashExample2() {
	// setup some filenames and content to be hasehed
	files := map[string]string{
		"file1.txt": "foo",
		"file2.txt": "foo",
		"file3.txt": "bar",
	}

	// create some files to hash
	for fn, cnt := range files {
		// create file
		file, err := os.Create(fn)
		if err != nil {
			panic("error")
		}
		// close file stream when function stops
		defer file.Close()
		// write content to file
		file.WriteString(cnt)
	}
	hash1, err := getHash("file1.txt")
	if err != nil {
		return
	}
	hash2, err := getHash("file2.txt")
	if err != nil {
		return
	}
	hash3, err := getHash("file3.txt")
	if err != nil {
		return
	}

	fmt.Println(hash1, hash2, hash3)
	fmt.Println(hash1 == hash2)
	fmt.Println(hash1 == hash3)
}

func cryptoHashExample() {
	// create hasher
	hasher := sha1.New()
	// create data
	data := []byte("very secret")
	hasher.Write(data)
	hashed := hasher.Sum([]byte{})
	fmt.Println(string(hashed))
}

func main() {
	hashExample1()
	hashExample2()
	cryptoHashExample()
}
