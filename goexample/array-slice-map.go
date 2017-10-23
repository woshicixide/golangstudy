package main

import (
	"fmt"
)

func main() {
	// var a [3]int
	// 数组可以在声明后直接使用
	// a[2] = 3

	// panic: runtime error: index out of range
	// 声明一个slice不可以直接赋值，报上述错误
	// var s []int
	// slice需要make后才能使用
	//  s := make([]int, 3)
	// s[1] = 1

	// 声明后直接使用报如下错误panic: assignment to entry in nil map
	// map也需要先make才能赋值
	// var m map[string]int
	m := make(map[string]int)
	m["s"] = 1

	fmt.Println(m)
}
