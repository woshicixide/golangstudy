package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// Buffered Channels
	// fatal error: all goroutines are asleep - deadlock!
	// c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
}
