package main

import (
	"fmt"
	ch "github.com/golang/groupcache/consistenthash"
)

func main() {
	c := ch.New(10, nil)
	c.Add("192.168.0.1","192.168.0.2","192.168.0.3","192.168.0.4","192.168.0.5","192.168.0.6")

	fmt.Println(c.Get("aa"))
	fmt.Println(c.Get("bb"))
	fmt.Println(c.Get("cc"))
	fmt.Println(c.Get("dd"))
	fmt.Println(c.Get("ee"))
	fmt.Println(c.Get("ff"))
	fmt.Println(c.Get("gg"))
	fmt.Println(c)
	}
