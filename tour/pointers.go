package main

import (
	"fmt"
)

// The type *T is a pointer to a T value. Its zero value is nil

// The & operator generates a pointer to its operand

// The * operator denotes the pointer's underlying value.

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
