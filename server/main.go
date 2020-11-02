package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8989"

func echoBack(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}

	fmt.Printf("received: %s", data)
	conn.Write([]byte(strings.ToUpper(data)))
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	fmt.Println("listening on %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accepting connection: %s\n", err.Error())
			continue
		}
		go echoBack(conn)
	}
}
