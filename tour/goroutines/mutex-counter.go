package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (self *SafeCounter) Inc(key string) {
	self.mux.Lock()
	self.v[key]++
	self.mux.Unlock()
}

func (self *SafeCounter) Value(key string) int {
	self.mux.Lock()
	defer self.mux.Unlock()
	return self.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
