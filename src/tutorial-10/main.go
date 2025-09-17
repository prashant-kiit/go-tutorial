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

func readMessage(msg string, t time.Duration, wg *sync.WaitGroup, m *sync.RWMutex) {
	for i := 1; i <= 1000000; i++ {
		m.RLock()
		fmt.Println(msg, i, count)
		time.Sleep(t * time.Millisecond)
		m.RUnlock()
	}
	wg.Done()
}

func foo1() {
	// Run as goroutine
	var wg sync.WaitGroup
	var m1 = sync.Mutex{}
	var m2 = sync.RWMutex{}
	wg.Add(4)
	go printMessage("Hello from Goroutine 1", 0, &wg, &m1)
	go printMessage("Hello from Goroutine 2", 0, &wg, &m1)
	go readMessage("Hello from Goroutine 3", 0, &wg, &m2)
	go readMessage("Hello from Goroutine 4", 0, &wg, &m2)
	wg.Wait()
	fmt.Println("All goroutines completed", count)
	// Main function also runs concurrently
	// printMessage("Hello from Main", 10)
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