package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var t I = T{"hello"}
	// var t T = T{"hello"}
	t.M()
}
