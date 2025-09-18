package main

import (
	"fmt"
)

func worker(messages chan string) {
	// send data into channel
	messages <- "Hello from worker"
}

func main() {
	// create a channel of type string
	messages := make(chan string)

	// start goroutine
	go worker(messages)

	// receive data from channel (blocking)
	msg := <-messages
	fmt.Println("Received:", msg)
}