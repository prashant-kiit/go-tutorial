package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("Worker %d: step %d", id, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)

	// start two workers
	go worker(1, ch)
	go worker(2, ch)

	// receive messages from both
	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}
}
