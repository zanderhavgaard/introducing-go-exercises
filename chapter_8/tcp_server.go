package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// create a tcp listener
	protocol := "tcp"
	port := ":8080"
	listener, err := net.Listen(protocol, port)
	if err != nil {
		fmt.Println(err)
		return
	}

	// infinite loop to keep receiving connections
	for {
		// accept a connection
		conn, err := listener.Accept()
		// if we encounter an error we simply continue listening
		// instead of crashing the server
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	// receive a message
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}
	conn.Close()
}

func client() {
	protocol := "tcp"
	host := "localhost:8080"
	// client to connect to the server
	conn, err := net.Dial(protocol, host)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello, world!"
	fmt.Println("Sending:", msg)

	// send message
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}

func main() {
	go server()
	go client()

	// hack to keep the program running while the server and client run in the background
	var input string
	fmt.Scanln(&input)
}
