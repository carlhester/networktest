package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func readInput(ch chan string) {
	defer close(ch)
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			close(ch)
			return
		}
		ch <- s
	}
}

func main() {
	ch := make(chan string)
	go readInput(ch)

stdinloop:
	for {
		select {
		case stdin, ok := <-ch:
			fmt.Println(stdin, ok)
			if !ok {
				break stdinloop
			} else {
				fmt.Println("Read input from stdin:", stdin)
			}
		case <-time.After(1 * time.Second):
			fmt.Println("waiting")
		}
	}
	fmt.Println("Done, stdin must be closed")
}
