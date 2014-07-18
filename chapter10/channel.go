package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; i < 3; i++ {
		c <- "ping"
	}
	c <- "eof"
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	messages := true
	for ; messages ; {
		msg := <- c
		if msg == "eof" {
			messages = false
		} else {
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}
	fmt.Println("done")
}
func anotherPrinter(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	go anotherPrinter(c)

	var input string
	fmt.Scanln(&input)
}
