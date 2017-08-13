package main

import (
	"fmt"
)

func fabonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		// fmt.Println(x, y, z)
		return z
	}
}

func main() {

	f := fabonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
