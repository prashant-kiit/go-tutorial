package main

import (
	"fmt"
	// "runtime"
	"time"
	"sync"
)

var count int = 0

func printMessage(msg string, t time.Duration, wg *sync.WaitGroup, m *sync.Mutex) {
	for i := 1; i <= 1000000; i++ {
		m.Lock()
		count++
		m.Unlock()
		fmt.Println(msg, i)
		time.Sleep(t * time.Millisecond)
	}

	wg.Done()
}

func foo1() {
	// Run as goroutine
	var wg sync.WaitGroup
	var m = sync.Mutex{}
	wg.Add(2)
	go printMessage("Hello from Goroutine 1", 0, &wg, &m)
	go printMessage("Hello from Goroutine 2", 0, &wg, &m)
	wg.Wait()

	// Main function also runs concurrently
	// printMessage("Hello from Main", 10)
	fmt.Println("All goroutines completed", count)
}

func main() {
	foo1()
	// fmt.Println("CPUs:", runtime.NumCPU())
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(2)) 
}

// Where do Goroutines fit?
// Goroutines are concurrent by default.
// The Go runtime (scheduler) multiplexes many goroutines onto a small number of OS threads.
// If you have multiple CPU cores and set GOMAXPROCS > 1 (default in modern Go), goroutines can also run in parallel.