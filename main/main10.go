package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is a demo
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc is a demo
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.v[key]++
}

// Value is a demo
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.v[key]
}

func main10() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 100000; i++ {
		go c.Inc("somekey1")
		go c.Inc("somekey2")
		go c.Inc("somekey3")
		go c.Inc("somekey4")
		go c.Inc("somekey5")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey1"))
	fmt.Println(c.Value("somekey2"))
	fmt.Println(c.Value("somekey3"))
	fmt.Println(c.Value("somekey4"))
	fmt.Println(c.Value("somekey5"))
}
