package main

import (
	// "fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	m := strings.Split(s, " ")

	for _, v := range m {
		_, t := r[v]
		if t {
			r[v]++
		} else {
			r[v] = 1
		}
	}
	// fmt.Println(r)
	return r
}

func main() {
	wc.Test(WordCount)
}
