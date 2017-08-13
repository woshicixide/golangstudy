package main

import (
	"fmt"
)

func do(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("Twice %d is %d\n", t, t*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", t, len(t))
	default:
		fmt.Printf("i do not konw about %T\n", t)
	}
}

func main() {
	do(42)
	do("hello")
	do(true)
}
