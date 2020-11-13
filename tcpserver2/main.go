package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var players []*player
	seq := 0
	cToS := make(chan string)
	addr := &net.TCPAddr{Port: 9998}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Printf("Listening on %+v\n", addr)
	for {
		time.Sleep(1 * time.Second)
		go acceptNew(seq, listener, players)
		seq += 1
		go interval(seq, cToS)
		go handleConn(players, cToS)
	}
}

func interval(seq int, cToS chan string) {
	for {
		fmt.Println(string(seq))
		cToS <- string(seq)
		time.Sleep(1 * time.Second)
	}
}

type player struct {
	*net.TCPConn
	seq int
}

func handleConn(players []*player, cToS chan string) {
	select {
	case msg := <-cToS:
		sendit(players, msg)
	default:
		return
	}
}

func sendit(players []*player, txt string) {
	for _, v := range players {
		fmt.Printf("sending %s to %s\n", txt, v.RemoteAddr())
		writer := bufio.NewWriter(v)
		_, _ = writer.WriteString(txt)
		writer.Flush()
	}
}

func acceptNew(seq int, listener *net.TCPListener, players []*player) {
	conn, err := listener.AcceptTCP()
	if err != nil {
		fmt.Println("error accepting connection: %s\n", err.Error())
	}
	fmt.Printf("client connected from %+v\n", conn.RemoteAddr())
	p := &player{
		seq:     seq,
		TCPConn: conn,
	}
	players = append(players, p)
}
