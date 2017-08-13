package main

import (
	"fmt"
)

func main() {
	// next line work
	// var sum int
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
