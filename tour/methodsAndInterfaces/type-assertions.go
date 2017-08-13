package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f)

	// If i does not hold a T, the statement will trigger a panic.
	f = i.(float64) // panic
	fmt.Println(f)

}
