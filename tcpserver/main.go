package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	cToS := make(chan string, 100)
	fmt.Println("1")
	addr := &net.TCPAddr{Port: 9998}
	listener, err := net.ListenTCP("tcp", addr)
	fmt.Println("2")
	if err != nil {
		panic(err)
	}
	fmt.Println("3")
	defer listener.Close()
	fmt.Printf("listening on %s\n", addr)
	fmt.Println("4")

	players := []*player{}
	fmt.Println("5")
	for {
		conn, err := listener.AcceptTCP()
		fmt.Println("6")
		if err != nil {
			fmt.Println("error accepting connection: %s\n", err.Error())
			continue
		}
		players = append(players, &player{conn})
		fmt.Println("7")
		cToS <- "test\n"

		fmt.Println("8")
		go handleConn(players, cToS)
	}
}

type player struct {
	*net.TCPConn
}

func handleConn(players []*player, cToS chan string) {
	var msg string
	fmt.Println("9")
	select {
	case msg := <-cToS:
		sendit(players, msg)
	default:
		fmt.Println(msg)
	}
}

func sendit(players []*player, txt string) {
	for _, v := range players {
		fmt.Println("10")
		fmt.Printf("Connect from: %s\n", v.RemoteAddr())
		fmt.Println("11")
		writer := bufio.NewWriter(v)
		fmt.Println("12")
		fmt.Println("13")

		fmt.Println("14")
		fmt.Println("15")
		fmt.Sprintf("%s\n", txt)
		_, _ = writer.WriteString(txt)
		writer.Flush()
	}
}
