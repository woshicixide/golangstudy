package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 1
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is a week day")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i'm  bool")
		case int:
			fmt.Println("i'm int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatIam(true)
	whatIam(1)
	whatIam("hay")

}
