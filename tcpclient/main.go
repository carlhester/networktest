package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:9998")
	if err != nil {
		panic(fmt.Sprintf("Unable to connect. %s", err.Error()))
	}
	reader := bufio.NewScanner(conn)
	for {
		time.Sleep(1 * time.Second)
		scan := reader.Scan()
		if err != nil {
			fmt.Printf("Err: %+v\n", err)
			continue
		}
		fmt.Printf("New Msgs?: %s\n", scan)
		if scan {
			txt := reader.Text()
			fmt.Println(txt)
		}
	}
	conn.Close()
}
