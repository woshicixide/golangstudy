package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(3, 54))
}
