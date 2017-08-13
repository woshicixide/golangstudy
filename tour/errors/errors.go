package main

import (
	"fmt"
	"time"
)

// The error type is a built-in interface similar to fmt.Stringer:

// type error interface {
//     Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (self *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", self.When, self.What)
}

func run() error {
	// without & get next error
	// cannot use MyError literal (type MyError) as type error in return argument:
	// MyError does not implement error (Error method has pointer receiver)
	return &MyError{time.Now(), "it do not work"}

}

func main() {
	if r := run(); nil != r {
		fmt.Println(r)
	}
}
