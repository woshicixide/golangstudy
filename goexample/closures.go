package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 闭包
// 两条链

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
