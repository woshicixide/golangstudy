package main

import (
	"fmt"
)

// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

// The loop for i := range c receives values from the channel repeatedly until it is closed.
func fibonacci(t int, c chan int) {
	x, y := 0, 1
	for i := 0; i < t; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
