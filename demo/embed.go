package main

import (
	"fmt"
)

func main() {
	ad := admin{user{name: "aa", email: "bb"}, "lll"}
	sayHello(ad.user)
	// ad.sayHello()
	// ad.user.sayHello()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

type Hello interface {
	hello()
}

func sayHello(u Hello) {
	u.hello()
}

func (u user) hello() {
	fmt.Println("i am a user")
}

// func (a admin) sayHello() {
// 	fmt.Println("i am a admin")
// }
