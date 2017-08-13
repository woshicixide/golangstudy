package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello,World")
	// fmt.Println(r.Len())

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v b=%v err=%v\n", n, b, err)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if io.EOF == err {
			break
		}
	}

}
