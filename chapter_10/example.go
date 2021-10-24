package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"http://apple.com",
		"http://amazon.com",
		"http://google.com",
		"http://microsoft.com",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		fmt.Println(result.URL, result.Size)
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest website:", biggest.URL, biggest.Size)
}
