package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for i := 0; i < 10; i++ {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var s string
	fmt.Scanln(&s)
}
