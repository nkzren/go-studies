package main

import (
	"fmt"
	"sync"
	"time"
)

// This should be safe
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map
	c.v[key]++
	c.mu.Unlock()
}

// Get returns current value of the counter for the given key
func (c *SafeCounter) Get(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Get("somekey"))
}
