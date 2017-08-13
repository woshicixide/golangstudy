package main

import (
	"fmt"
)

func main() {
	sum := 1
	// formated the next line
	// for ;sum < 1000; {
	// like while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
