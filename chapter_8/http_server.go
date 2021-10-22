package main

import (
	"io"
	"net/http"
)

// setup handler function for the 'hello' endpoint
func hello(res http.ResponseWriter, req *http.Request) {
	// set headers for response
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	body := `<DOCTYPE html>
	<html>
		<head>
			<title>Hello world!</title>
		</head>
		<body>
			<h1>Hello from a golang http server!</h1>
		</body>
	</html>`

	io.WriteString(res, body)
}

func main() {
	// setup the endpoint handler
	http.HandleFunc("/hello", hello)
	// start the http server
	http.ListenAndServe(":8080", nil)
}
