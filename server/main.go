package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8989"

func main() {
	// listen for a connection
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	fmt.Printf("listening on %s\n", addr)
	// loop forever
	for {
		// accept a connection when it shows up
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accepting connection: %s\n", err.Error())
			continue
		}
		// pass connection to goroutine
		go echoBack(conn)
	}
}

func echoBack(conn net.Conn) {

	// used to read from connection
	reader := bufio.NewReader(conn)

	// assign received data to data
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}

	fmt.Printf("received: %s", data)
	// send response by writing to conn
	fmt.Printf("sending %s\n", strings.ToUpper(data))
	conn.Write([]byte(strings.ToUpper(data)))
	fmt.Printf("sending Welcome\n")
	conn.Write([]byte("Welcome!\n"))
	defer conn.Close()
}
