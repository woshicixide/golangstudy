package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Answer"] = 44
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// If key is not in the map, then elem is the zero value for the map's element type.
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
