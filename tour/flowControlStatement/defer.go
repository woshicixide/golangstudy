package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("wrold")
	fmt.Println("hello")
}
