package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.RWMutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()         // full write lock
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Get(key string) int {
	c.mu.RLock()        // read lock
	defer c.mu.RUnlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	// Writer goroutine
	go func() {
		for i := 0; i < 5; i++ {
			c.Inc("prashant")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Multiple reader goroutines
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				fmt.Printf("Reader %d: %d\n", id, c.Get("prashant"))
				time.Sleep(150 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
}
