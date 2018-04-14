package main

import (
	"fmt"
	"net/http"
)

const MaxOutstanding = 100

var sem = make(chan int, MaxOutstanding)

func init() {
	for i := 0; i < MaxOutstanding; i++ {
		sem <- 1
	}
}

func main() {
	server()
}

// func server(seqs chan *http.Request) {
// 	for {
// 		s := <-seq
// 		go handle(s)
// 	}
// }

// func handle(seq *http.Request) {
// 	// 如果资源耗尽，会被阻塞
// 	<-sem
// 	process(seq)
// 	sem <- 1
// }

func server(seqs chan *http.Resquest) {
	for seq := range seq {
		<-sem
		go func() {
			process(seq)
			sem <- 1
		}()
	}
}

func process(seq *http.Resquest) {
	fmt.Println(seq)
}
