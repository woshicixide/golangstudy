package main

import (
	"fmt"
)

//The type []T is a slice with elements of type T

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
