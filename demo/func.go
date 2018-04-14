package main

import "fmt"

var (
	portPicker func(a string) string
)


func main() {
	//打印出了nil
	fmt.Println(portPicker)
}
