package main

import (
	"fmt"
)

func main() {
	var a []int
	printSlice(a)

	s := append(a, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
	// len=5 cap=6 [0 1 2 3 4]
	// cap is 6
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
