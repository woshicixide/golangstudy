package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {
	t := tree.New(1)
	ch := make(chan int)

	go walk(t, ch)

	for i := range ch {
		fmt.Println(i)
	}
	s := tree.New(2)
	fmt.Println(sum(s, t))

}

func sum(s, t *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go walk(s, ch1)
	go walk(t, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if nil == t {
		return
	}

	_walk(t.Left, ch)
	ch <- t.Value
	_walk(t.Right, ch)
}
