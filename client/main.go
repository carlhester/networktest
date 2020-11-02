package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8989"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter text: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading string: %s\n", err.Error())
			continue
		}

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("error with dial: %s\n", err.Error())
			continue
		}

		fmt.Fprintf(conn, data)

		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error with reading response: %s\n", err.Error())
			continue
		}

		fmt.Printf("Received back: %s\n", status)
		conn.Close()
	}
}
