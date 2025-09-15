package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMessage(msg string, t time.Duration) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(t * time.Millisecond)
	}
}

func foo1() {
	// Run as goroutine
	go printMessage("Hello from Goroutine 1", 5000)
	go printMessage("Hello from Goroutine 2", 10000)

	// Main function also runs concurrently
	printMessage("Hello from Main", 10)
}

func main() {
	foo1()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(2)) 
}

// Where do Goroutines fit?
// Goroutines are concurrent by default.
// The Go runtime (scheduler) multiplexes many goroutines onto a small number of OS threads.
// If you have multiple CPU cores and set GOMAXPROCS > 1 (default in modern Go), goroutines can also run in parallel.